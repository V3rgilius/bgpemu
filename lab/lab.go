package lab

import (
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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
		log.Infof("Trying in-cluster configuration")
		rCfg, err := rest.InClusterConfig()
		if err != nil {
			log.Infof("Falling back to kubeconfig: %q", m.kubecfg)
			rCfg, err = clientcmd.BuildConfigFromFlags("", m.kubecfg)
			if err != nil {
				return nil, err
			}
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

func (m *Manager) GetGrpcServers(pods []string, topo string) error {
	if m.gServers == nil {
		m.gServers = make(map[string]string, 100)
	}
	for _, pod := range pods {
		if m.gServers[pod] == "" {
			continue
		}
		svcName := fmt.Sprintf("service-%s", pod)
		svc, err := m.kClient.CoreV1().Services(topo).Get(context.Background(), svcName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		for _, port := range svc.Spec.Ports {
			if port.Port == 50051 {
				m.gServers[pod] = fmt.Sprintf("%s:%d", svc.Spec.ExternalIPs[0], port.Port)
				break
			}
		}
	}
	return nil
}

func (m *Manager) GetGrpcServersAll(topo string) error {
	m.gServers = make(map[string]string, 100)
	svcList, err := m.kClient.CoreV1().Services(topo).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, svc := range svcList.Items {
		for _, port := range svc.Spec.Ports {
			if port.Port == 50051 {
				m.gServers[strings.Split(svc.Name, "-")[1]] = fmt.Sprintf("%s:%d", svc.Spec.ExternalIPs[0], port.Port)
				break
			}
		}
	}
	return nil
}

func (m *Manager) GetGServers() map[string]string {
	return m.gServers
}
