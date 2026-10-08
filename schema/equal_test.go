package schema

import (
	"encoding/json/v2"
	"testing"

	"github.com/MarkRosemaker/openapi"
)

// TestEqualAndSameShape compares the two properties of each case in testdata/cases.json, either way round.
func TestEqualAndSameShape(t *testing.T) {
	t.Parallel()

	doc, err := openapi.LoadFromFile("testdata/cases.json")
	if err != nil {
		t.Fatal(err)
	}

	for name, c := range doc.Components.Schemas.ByIndex() {
		if c.Extensions == nil {
			continue // a schema the cases refer to
		}

		var want struct {
			Equal     bool `json:"x-equal"`
			SameShape bool `json:"x-same-shape"`
		}
		if err := json.Unmarshal(c.Extensions, &want); err != nil {
			t.Fatalf("%s: %v", name, err)
		}

		a, b := c.Properties["a"], c.Properties["b"]
		if a == nil || b == nil {
			t.Fatalf("%s: a case compares its properties a and b", name)
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for _, pair := range [][2]*openapi.Schema{{a, b}, {b, a}} {
				if got := Equal(pair[0], pair[1]); got != want.Equal {
					t.Errorf("Equal = %v, want %v", got, want.Equal)
				}

				if got := SameShape(pair[0], pair[1]); got != want.SameShape {
					t.Errorf("SameShape = %v, want %v", got, want.SameShape)
				}
			}
		})
	}
}

// TestEqualAndSameShape_Nil: no schema is only the same as no schema, which a document cannot hold.
func TestEqualAndSameShape_Nil(t *testing.T) {
	t.Parallel()

	s := &openapi.Schema{Type: openapi.TypeString}

	if !Equal(nil, nil) || !SameShape(nil, nil) {
		t.Error("nil differs from nil")
	}

	if Equal(s, nil) || Equal(nil, s) || SameShape(s, nil) || SameShape(nil, s) {
		t.Error("a schema is the same as none")
	}
}
