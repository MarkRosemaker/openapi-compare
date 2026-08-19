package schema

import (
	"testing"

	"github.com/MarkRosemaker/openapi"
)

func inlineRef(s *openapi.Schema) *openapi.SchemaRef {
	return &openapi.SchemaRef{Value: s}
}

func namedRef(id, summary, description string) *openapi.SchemaRef {
	return &openapi.SchemaRef{Ref: &openapi.Reference{
		Identifier:  id,
		Summary:     summary,
		Description: description,
	}}
}

func stringSchema() *openapi.Schema {
	return &openapi.Schema{Type: openapi.TypeString}
}

func TestEqualAndSameShape(t *testing.T) {
	tests := []struct {
		name          string
		a, b          *openapi.Schema
		wantEqual     bool
		wantSameShape bool
	}{
		{
			name:          "both nil",
			a:             nil,
			b:             nil,
			wantEqual:     true,
			wantSameShape: true,
		},
		{
			name:          "one nil",
			a:             stringSchema(),
			b:             nil,
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "identical",
			a:             stringSchema(),
			b:             stringSchema(),
			wantEqual:     true,
			wantSameShape: true,
		},
		{
			name:          "different type",
			a:             &openapi.Schema{Type: openapi.TypeString},
			b:             &openapi.Schema{Type: openapi.TypeInteger},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different title only",
			a:             &openapi.Schema{Type: openapi.TypeString, Title: "A"},
			b:             &openapi.Schema{Type: openapi.TypeString, Title: "B"},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "different description only",
			a:             &openapi.Schema{Type: openapi.TypeString, Description: "one"},
			b:             &openapi.Schema{Type: openapi.TypeString, Description: "two"},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "different default only",
			a:             &openapi.Schema{Type: openapi.TypeString, Default: []byte(`"a"`)},
			b:             &openapi.Schema{Type: openapi.TypeString, Default: []byte(`"b"`)},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "different example only",
			a:             &openapi.Schema{Type: openapi.TypeString, Example: []byte(`"a"`)},
			b:             &openapi.Schema{Type: openapi.TypeString, Example: []byte(`"b"`)},
			wantEqual:     true,
			wantSameShape: true,
		},
		{
			name:          "different extensions",
			a:             &openapi.Schema{Type: openapi.TypeString, Extensions: []byte(`{"x-a":1}`)},
			b:             &openapi.Schema{Type: openapi.TypeString, Extensions: []byte(`{"x-a":2}`)},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name: "different oneOf",
			a: &openapi.Schema{OneOf: openapi.SchemaRefList{
				inlineRef(&openapi.Schema{Type: openapi.TypeString}),
			}},
			b: &openapi.Schema{OneOf: openapi.SchemaRefList{
				inlineRef(&openapi.Schema{Type: openapi.TypeInteger}),
			}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name: "different anyOf",
			a: &openapi.Schema{AnyOf: openapi.SchemaRefList{
				inlineRef(&openapi.Schema{Type: openapi.TypeString}),
			}},
			b: &openapi.Schema{AnyOf: openapi.SchemaRefList{
				inlineRef(&openapi.Schema{Type: openapi.TypeBoolean}),
			}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different not",
			a:             &openapi.Schema{Not: inlineRef(&openapi.Schema{Type: openapi.TypeString})},
			b:             &openapi.Schema{Not: inlineRef(&openapi.Schema{Type: openapi.TypeInteger})},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name: "nested property differs only in description",
			a: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.SchemaRefs{
					"name": inlineRef(&openapi.Schema{Type: openapi.TypeString, Description: "one"}),
				},
			},
			b: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.SchemaRefs{
					"name": inlineRef(&openapi.Schema{Type: openapi.TypeString, Description: "two"}),
				},
			},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name: "nested property differs in type",
			a: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.SchemaRefs{
					"name": inlineRef(&openapi.Schema{Type: openapi.TypeString}),
				},
			},
			b: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.SchemaRefs{
					"name": inlineRef(&openapi.Schema{Type: openapi.TypeInteger}),
				},
			},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "same ref identifier, different summary",
			a:             &openapi.Schema{Properties: openapi.SchemaRefs{"p": namedRef("#/components/schemas/Foo", "a", "")}},
			b:             &openapi.Schema{Properties: openapi.SchemaRefs{"p": namedRef("#/components/schemas/Foo", "b", "")}},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "different ref identifier",
			a:             &openapi.Schema{Properties: openapi.SchemaRefs{"p": namedRef("#/components/schemas/Foo", "", "")}},
			b:             &openapi.Schema{Properties: openapi.SchemaRefs{"p": namedRef("#/components/schemas/Bar", "", "")}},
			wantEqual:     false,
			wantSameShape: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.a, tt.b); got != tt.wantEqual {
				t.Errorf("Equal() = %v, want %v", got, tt.wantEqual)
			}
			if got := Equal(tt.b, tt.a); got != tt.wantEqual {
				t.Errorf("Equal() (swapped) = %v, want %v", got, tt.wantEqual)
			}
			if got := SameShape(tt.a, tt.b); got != tt.wantSameShape {
				t.Errorf("SameShape() = %v, want %v", got, tt.wantSameShape)
			}
			if got := SameShape(tt.b, tt.a); got != tt.wantSameShape {
				t.Errorf("SameShape() (swapped) = %v, want %v", got, tt.wantSameShape)
			}
		})
	}
}
