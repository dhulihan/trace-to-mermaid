package ttm

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Options struct {
	// Depth controls how many levels of nesting we want to maintain
	// from the trace to the diagram.
	depth int

	// Direction can be TD, LR, etc
	direction string
}

func (o *Options) Direction() string {
	if o == nil || o.direction == "" {
		return "TD"
	}

	return o.direction
}

// TODO: replace this with the true struct
type JaegerTrace struct {
	Data []struct {
		Spans []struct {
			OperationName string
		}
		// Processes are the services involved
		Processes map[string]struct {
			ServiceName string
			Tags        []struct {
				Key   string
				Type  string
				Value string
			}
		}
	}
}

func ParseJaegerTrace(b []byte) (*JaegerTrace, error) {
	t := &JaegerTrace{}

	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("could not parse jaeger trace: %w", err)
	}

	if len(t.Data) == 0 {
		return nil, fmt.Errorf("data is empty")
	}

	return t, nil
}

func (t *JaegerTrace) ToMermaidFlowDiagram(opts *Options) (string, error) {
	// use this for now, we could introduce a more sophisticated
	// struct to contain a flow diagram
	var b strings.Builder

	b.WriteString("flowchart " + opts.Direction())
	p, err := t.processes(opts)
	if err != nil {
		return "", fmt.Errorf("could not render processes: %w", err)
	}
	b.WriteString(p)

	// processes only for now

	return b.String(), nil
}

func (t *JaegerTrace) processes(opts *Options) (string, error) {
	if len(t.Data[0].Processes) == 0 {
		return "", fmt.Errorf("no processes found in trace")
	}

	// use this for now, we could introduce a more sophisticated
	// struct to contain a flow diagram
	var b strings.Builder

	for _, p := range t.Data[0].Processes {
		b.WriteString(p.ServiceName)
	}

	return b.String(), nil
}
