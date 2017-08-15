package config

import (
	"os"
	"path/filepath"
)

var workingDir, _ = os.Getwd()
var templatesDir = filepath.Join(workingDir, "webapp", "templates")
