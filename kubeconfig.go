package kubeconfig

import (
	"bitbucket.org/linkernetworks/cv-tracker/src/env"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func FindConfig() (string, bool) {
	if home := env.HomeDir(); home != "" {
		p := filepath.Join(home, ".kube", "config")
		_, err := os.Stat(p)
		if err != nil {
			return "", false
		}
		return p, true
	}
	return "", false
}

func Load(context string, kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		_, err := os.Stat(kubeconfig)
		if err == nil {
			return clientcmd.BuildConfigFromFlags(context, kubeconfig)
		}
	}
	mykubeconfig, found := FindConfig()
	if found {
		return clientcmd.BuildConfigFromFlags(context, mykubeconfig)
	}
	return rest.InClusterConfig()
}
