package application

import (
	"log"
	"os"
	"path/filepath"
)

type ApplicationConfig struct {
	KubeConfigPath string
	TemplatesDir   string
}

func NewAppConfig() *ApplicationConfig {
	kubeConfigPath := os.Getenv("KUBE_CONFIG")
	if kubeConfigPath == "" {
		log.Fatal("KUBE_CONFIG evironment variable must be set")
	}
	workingDir, err := os.Getwd()
	if err != nil {
		log.Panic(err.Error())
	}
	templatesDir := filepath.Join(workingDir, "webapp", "templates")

	return &ApplicationConfig{kubeConfigPath, templatesDir}
}
