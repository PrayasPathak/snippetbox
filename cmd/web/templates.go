package main

import "github.com/PrayasPathak/snippetbox/internal/models"

type TemplateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
