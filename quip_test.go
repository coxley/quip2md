package quip2md

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarkownConversion(t *testing.T) {
	e, err := AssetString("assets/expected.md")
	if err != nil {
		t.Errorf("Error raised when fetching asset:\n%v", err)
	}
	html, err := AssetString("assets/example.html")
	if err != nil {
		t.Errorf("Error raised when fetching asset:\n%v", err)
	}
	a, err := QuipToMarkdown(html)
	if err != nil {
		t.Errorf("Error raised when converting to Markdown:\n%v", err)
	}

	if diff := cmp.Diff(e, a); diff != "" {
		t.Errorf("Expected vs conversion: differs: (-want +got)\n%s", diff)
	}

}
