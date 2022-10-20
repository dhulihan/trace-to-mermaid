package mermaid

import (
	"bytes"
	"fmt"
)

type Connection struct {
	Src []Node
	Dst []Node

	Style      string
	Terminator string
}

type Node struct {
	Name string

	Connections []Connection
}

type Flowchart struct {
	Direction string
	Nodes     map[string]string
}

func (fc *Flowchart) Render() (string, error) {
	if flowchartTemplate == nil {
		return "", fmt.Errorf("flowchart template not loaded")
	}

	bb := new(bytes.Buffer)
	if err := flowchartTemplate.Execute(bb, fc); err != nil {
		return "", fmt.Errorf("could not execute template: %w", err)
	}
	return string(bb.Bytes()), nil
}
