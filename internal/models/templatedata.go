package models

import "github.com/FeiBoNaQi/bookings/internal/forms"

// TemplateData holds data send from handler to
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
