package tabwriter

import (
	"bytes"
	"testing"
)

func TestPrettyWrite(t *testing.T) {
	var output bytes.Buffer

	writable := "Bleus P1"
	PrettyWrite(writable, &output)
	want := "Bleus P1\n"
	got := output.String()

	if output.String() != want {
		t.Errorf("Expected '%s' but got '%s'", want, got)
	}
}
