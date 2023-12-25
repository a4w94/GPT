package prompt

import (
	"embed"
	"gpt/util"
	"log"
)

//go:embed templates/*
var templatesFS embed.FS
var templatesDir = "templates"

// Template file names
const (
	RefactorTemplate = "refactor.tmpl"
)

// Initializes the prompt package by loading the templates from the embedded file system.
func init() {
	if err := util.LoadTemplates(templatesFS, templatesDir); err != nil {
		log.Fatal(err)
	}
}
