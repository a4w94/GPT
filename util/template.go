package util

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
)

// custom type for the template data.
type Data map[string]interface{}

var (
	templates map[string]*template.Template
	//templatesDir = "templates"
)

func LoadTemplates(files embed.FS, templatesDir string) error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			return err
		}

		templates[tmpl.Name()] = pt
	}
	return nil
}

// GetTemplateByString returns the parsed template as a string.
func GetTemplateByString(name string, data map[string]interface{}) (string, error) {
	tpl, err := processTemplate(name, data)
	return tpl.String(), err
}

// processTemplate processes the template with the given name and data.
func processTemplate(name string, data map[string]interface{}) (*bytes.Buffer, error) {
	t, ok := templates[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}

	var tpl bytes.Buffer

	if err := t.Execute(&tpl, data); err != nil {
		return nil, err
	}

	return &tpl, nil
}
