package ttm

import (
	"encoding/json"
	"fmt"

	"github.com/dhulihan/trace-to-mermaid/pkg/mermaid"
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

func (t *JaegerTrace) ToMermaidFlowchart(opts *Options) (string, error) {
	flowchart := &mermaid.Flowchart{
		Direction: opts.Direction(),
	}

	// processes only for now
	return flowchart.Render()
}
