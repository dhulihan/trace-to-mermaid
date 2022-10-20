package mermaid

import (
	_ "embed"
	"fmt"
	"text/template"
)

var (
	//go:embed templates/flowchart.mmd
	flowchartTemplateRaw string
	flowchartTemplate    *template.Template
)

func LoadTemplates() error {
	if flowchartTemplateRaw == "" {
		return fmt.Errorf("flowchart template contents are empty")
	}
	// load flowchart
	var err error
	flowchartTemplate, err = template.New("flowchart").Parse(flowchartTemplateRaw)
	if err != nil {
		return fmt.Errorf("could not parse flowchart template: %w", err)
	}

	return nil
}
