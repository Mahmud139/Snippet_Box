package main

import "github.com/mahmud139/Snippet_Box/pkg/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}