package helper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

func DefaultKubeCfg() string {
	if v := os.Getenv("KUBECONFIG"); v != "" {
		return v
	}
	if home := homedir.HomeDir(); home != "" {
		return filepath.Join(home, ".kube", "config")
	}
	return ""
}

func FileRelative(p string) (string, error) {
	bp, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	return filepath.Dir(bp), nil
}

func GetTopologyTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Topology",
		APIVersion: "networkop.co.uk/v1beta1",
	}
}
