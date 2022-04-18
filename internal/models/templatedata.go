package models

import "github.com/TYHXX/go-miniBookingSystem/internal/forms"

// TempateData holds data sent from handlers tp templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
