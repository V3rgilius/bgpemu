package lab

import (
	"context"
	"fmt"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/v3rgilius/bgpemu/helper"
	"google.golang.org/protobuf/encoding/protojson"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

type Manager struct {
	kubecfg string
	kClient kubernetes.Interface
	rCfg    *rest.Config
	// basePath string
	topo     string
	gServers map[string]string
}

var protojsonUnmarshaller = protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: false,
}

func New(topo string) (*Manager, error) {
	m := &Manager{topo: topo}
	if m.rCfg == nil {
		// log.Infof("Trying in-cluster configuration")
		// rCfg, err := rest.InClusterConfig()
		m.kubecfg = helper.DefaultKubeCfg()
		log.Infof("Falling back to kubeconfig: %q", m.kubecfg)
		rCfg, err := clientcmd.BuildConfigFromFlags("", m.kubecfg)
		if err != nil {
			return nil, err
		}
		m.rCfg = rCfg
	}
	if m.kClient == nil {
		kClient, err := kubernetes.NewForConfig(m.rCfg)
		if err != nil {
			return nil, err
		}
		m.kClient = kClient
	}
	return m, nil
}

func (m *Manager) GetGrpcServers(pods []string) error {
	if m.gServers == nil {
		m.gServers = make(map[string]string, 100)
	}
	for _, pod := range pods {
		if m.gServers[pod] == "" {
			continue
		}
		svcName := fmt.Sprintf("service-%s", pod)
		svc, err := m.kClient.CoreV1().Services(m.topo).Get(context.Background(), svcName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		for _, port := range svc.Spec.Ports {
			if port.Port == 50051 {
				m.gServers[pod] = fmt.Sprintf("%s:%d", svc.Status.LoadBalancer.Ingress[0].IP, port.Port)
				break
			}
		}
	}
	return nil
}

func (m *Manager) GetGrpcServersAll() error {
	m.gServers = make(map[string]string, 100)
	svcList, err := m.kClient.CoreV1().Services(m.topo).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, svc := range svcList.Items {
		for _, port := range svc.Spec.Ports {
			if port.Port == 50051 {
				ingresses := svc.Status.LoadBalancer.Ingress
				// log.Infof("%s - %s", svc.Name, ingresses[0].IP)
				m.gServers[strings.Split(svc.Name, "-")[1]] = fmt.Sprintf("%s:%d", ingresses[0].IP, port.Port)
				break
			}

		}
	}
	return nil
}

func (m *Manager) GetGServers() map[string]string {
	return m.gServers
}

func (m *Manager) Exec(ctx context.Context, cmds []string, podname string, name string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	for _, command := range cmds {
		cmd := []string{
			"/bin/sh",
			"-c",
			command,
		}
		req := m.kClient.CoreV1().RESTClient().Post().
			Resource("pods").
			Name(podname).
			Namespace(m.topo).
			SubResource("exec")
		req.VersionedParams(&corev1.PodExecOptions{
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			Container: name,
			Command:   cmd,
			TTY:       false,
		}, scheme.ParameterCodec)
		log.Infof("Executing extra commands on container %s: %s", name, command)
		exec, err := remotecommand.NewSPDYExecutor(m.rCfg, "POST", req.URL())
		if err != nil {
			log.Errorf("error in creating executor for extra commands of container %s : %s", name, err.Error())
			return err
		}
		err = exec.Stream(remotecommand.StreamOptions{
			Stdin:  stdin,
			Stdout: stdout,
			Stderr: stderr,
			Tty:    false,
		})
		if err != nil {
			log.Errorf("error in executing extra commands of node %s : %s", name, err.Error())
			return err
		}
	}
	return nil
}
