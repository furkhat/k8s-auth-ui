package config

import (
	"html/template"
	"path/filepath"
)

var Templates = map[string]*template.Template{
	"serviceaccounts_list": template.Must(
		template.ParseFiles(
			filepath.Join(templatesDir, "serviceaccounts_list.html"),
		),
	),
}
