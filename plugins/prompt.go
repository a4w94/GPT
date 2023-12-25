package plugins

import (
	"embed"
	"gpt/util"
	"log"
)

//go:embed **/*
var templatesFS embed.FS

// Initializes the prompt package by loading the templates from the embedded file system.
func InitTemplate(templatesDir string) {
	if err := util.LoadTemplates(templatesFS, templatesDir); err != nil {
		log.Println("Error loading templates from embedded file system:", err)
		log.Fatal(err)
	}
}
