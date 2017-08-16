package application

import (
	"log"
	"os"
	"path/filepath"
)

type ApplicationConfig struct {
	TemplatesDir   string
}

func NewAppConfig() *ApplicationConfig {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Panic(err.Error())
	}
	templatesDir := filepath.Join(workingDir, "webapp", "templates")

	return &ApplicationConfig{templatesDir}
}
