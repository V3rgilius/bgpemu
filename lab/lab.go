package lab

import (
	log "github.com/sirupsen/logrus"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Manager struct {
	kubecfg  string
	kClient  kubernetes.Interface
	rCfg     *rest.Config
	basePath string
	topo     string
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
