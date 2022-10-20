package ttm

import (
	"io/ioutil"
	"os"
	"testing"
)

func loadTrace(t *testing.T, filename string) *JaegerTrace {
	t.Helper()
	goldenPath := "../../test/" + filename

	f, err := os.OpenFile(goldenPath, os.O_RDWR, 0644)
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	trace, err := ParseJaegerTrace(content)
	if err != nil {
		t.Fatalf("Error parsing trace %s: %s", goldenPath, err)
	}

	return trace
}

func TestJaegerTrace_ToMermaidFlowDiagram(t *testing.T) {
	type args struct {
		opts *Options
	}
	tests := []struct {
		name    string
		trace   *JaegerTrace
		args    args
		want    string
		wantErr bool
	}{
		{
			name:  "happy path",
			trace: loadTrace(t, "example-trace.json"),
			want: `
			flowchart TD
			foo
			`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.trace.ToMermaidFlowDiagram(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("JaegerTrace.ToMermaidFlowDiagram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JaegerTrace.ToMermaidFlowDiagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
