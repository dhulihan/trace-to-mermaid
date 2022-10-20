package mermaid

type FlowChart struct {
	Nodes []Node
}

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
