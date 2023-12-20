package prompt

import (
	"embed"
	"gpt/util"
	"log"
)

//go:embed templates/*
var templatesFS embed.FS

// Template file names
const (
	TranslationTemplate = "translation.tmpl"
	RefactorTemplate    = "refactor.tmpl"
)

// Initializes the prompt package by loading the templates from the embedded file system.
func init() {
	if err := util.LoadTemplates(templatesFS); err != nil {
		log.Fatal(err)
	}
}
