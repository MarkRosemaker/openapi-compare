package schema

import (
	"encoding/json/jsontext"
	"testing"

	"github.com/MarkRosemaker/openapi"
)

func ref(id, description string) *openapi.Schema {
	return &openapi.Schema{Description: description, Ref: &openapi.SchemaRef{Identifier: id}}
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
			a: &openapi.Schema{OneOf: openapi.SchemaList{
				&openapi.Schema{Type: openapi.TypeString},
			}},
			b: &openapi.Schema{OneOf: openapi.SchemaList{
				&openapi.Schema{Type: openapi.TypeInteger},
			}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name: "different anyOf",
			a: &openapi.Schema{AnyOf: openapi.SchemaList{
				&openapi.Schema{Type: openapi.TypeString},
			}},
			b: &openapi.Schema{AnyOf: openapi.SchemaList{
				&openapi.Schema{Type: openapi.TypeBoolean},
			}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different not",
			a:             &openapi.Schema{Not: &openapi.Schema{Type: openapi.TypeString}},
			b:             &openapi.Schema{Not: &openapi.Schema{Type: openapi.TypeInteger}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name: "nested property differs only in description",
			a: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.Schemas{
					"name": &openapi.Schema{Type: openapi.TypeString, Description: "one"},
				},
			},
			b: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.Schemas{
					"name": &openapi.Schema{Type: openapi.TypeString, Description: "two"},
				},
			},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name: "nested property differs in type",
			a: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.Schemas{
					"name": &openapi.Schema{Type: openapi.TypeString},
				},
			},
			b: &openapi.Schema{
				Type: openapi.TypeObject,
				Properties: openapi.Schemas{
					"name": &openapi.Schema{Type: openapi.TypeInteger},
				},
			},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "same ref, different description beside it",
			a:             &openapi.Schema{Properties: openapi.Schemas{"p": ref("#/components/schemas/Foo", "a")}},
			b:             &openapi.Schema{Properties: openapi.Schemas{"p": ref("#/components/schemas/Foo", "b")}},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "different ref identifier",
			a:             &openapi.Schema{Properties: openapi.Schemas{"p": ref("#/components/schemas/Foo", "")}},
			b:             &openapi.Schema{Properties: openapi.Schemas{"p": ref("#/components/schemas/Bar", "")}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different nullable",
			a:             &openapi.Schema{Type: openapi.TypeString, Nullable: true},
			b:             stringSchema(),
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different const",
			a:             &openapi.Schema{Const: jsontext.Value(`400`)},
			b:             &openapi.Schema{Const: jsontext.Value(`401`)},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different prefixItems",
			a:             &openapi.Schema{Type: openapi.TypeArray, PrefixItems: openapi.SchemaList{stringSchema()}},
			b:             &openapi.Schema{Type: openapi.TypeArray, PrefixItems: openapi.SchemaList{&openapi.Schema{Type: openapi.TypeInteger}}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "additionalProperties absent and true",
			a:             &openapi.Schema{Type: openapi.TypeObject},
			b:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Allowed: true}},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "additionalProperties true and the empty schema",
			a:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Allowed: true}},
			b:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Schema: &openapi.Schema{Description: "anything"}}},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "additionalProperties true and false",
			a:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Allowed: true}},
			b:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "additionalProperties absent and false",
			a:             &openapi.Schema{Type: openapi.TypeObject},
			b:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "additionalProperties boolean and schema",
			a:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{}},
			b:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Schema: stringSchema()}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "additionalProperties schemas differ only in description",
			a:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Schema: &openapi.Schema{Type: openapi.TypeString, Description: "a"}}},
			b:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Schema: &openapi.Schema{Type: openapi.TypeString, Description: "b"}}},
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "additionalProperties identical",
			a:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Schema: stringSchema()}},
			b:             &openapi.Schema{Type: openapi.TypeObject, AdditionalProperties: &openapi.AdditionalProperties{Schema: stringSchema()}},
			wantEqual:     true,
			wantSameShape: true,
		},
		{
			name:          "ref and inline",
			a:             &openapi.Schema{Properties: openapi.Schemas{"p": ref("#/components/schemas/Foo", "")}},
			b:             &openapi.Schema{Properties: openapi.Schemas{"p": stringSchema()}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "ref deprecated beside it",
			a:             &openapi.Schema{Deprecated: true, Ref: &openapi.SchemaRef{Identifier: "#/components/schemas/Foo"}},
			b:             ref("#/components/schemas/Foo", ""),
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "ref with a constraint beside it",
			a:             &openapi.Schema{Type: openapi.TypeString, MaxLength: new(uint(3)), Ref: &openapi.SchemaRef{Identifier: "#/components/schemas/Foo"}},
			b:             &openapi.Schema{Type: openapi.TypeString, Ref: &openapi.SchemaRef{Identifier: "#/components/schemas/Foo"}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different minLength",
			a:             &openapi.Schema{Type: openapi.TypeString, MinLength: 1},
			b:             stringSchema(),
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different maxLength",
			a:             &openapi.Schema{Type: openapi.TypeString, MaxLength: new(uint(5))},
			b:             &openapi.Schema{Type: openapi.TypeString, MaxLength: new(uint(6))},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different exclusiveMinimum",
			a:             &openapi.Schema{Type: openapi.TypeNumber, ExclusiveMin: new(0.0)},
			b:             &openapi.Schema{Type: openapi.TypeNumber},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different exclusiveMaximum",
			a:             &openapi.Schema{Type: openapi.TypeNumber, ExclusiveMax: new(1.0)},
			b:             &openapi.Schema{Type: openapi.TypeNumber, ExclusiveMax: new(2.0)},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different uniqueItems",
			a:             &openapi.Schema{Type: openapi.TypeArray, UniqueItems: true},
			b:             &openapi.Schema{Type: openapi.TypeArray},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different maxProperties",
			a:             &openapi.Schema{Type: openapi.TypeObject, MaxProperties: new(uint(1))},
			b:             &openapi.Schema{Type: openapi.TypeObject},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different discriminator",
			a:             &openapi.Schema{OneOf: openapi.SchemaList{stringSchema()}, Discriminator: &openapi.Discriminator{PropertyName: "kind"}},
			b:             &openapi.Schema{OneOf: openapi.SchemaList{stringSchema()}, Discriminator: &openapi.Discriminator{PropertyName: "type"}},
			wantEqual:     false,
			wantSameShape: false,
		},
		{
			name:          "different deprecated only",
			a:             &openapi.Schema{Type: openapi.TypeString, Deprecated: true},
			b:             stringSchema(),
			wantEqual:     false,
			wantSameShape: true,
		},
		{
			name:          "different examples only",
			a:             &openapi.Schema{Type: openapi.TypeString, Examples: []jsontext.Value{jsontext.Value(`"a"`)}},
			b:             stringSchema(),
			wantEqual:     true,
			wantSameShape: true,
		},
		{
			name:          "self-referential, not followed",
			a:             &openapi.Schema{Type: openapi.TypeObject, Properties: openapi.Schemas{"next": ref("#/components/schemas/Node", "")}},
			b:             &openapi.Schema{Type: openapi.TypeObject, Properties: openapi.Schemas{"next": ref("#/components/schemas/Node", "")}},
			wantEqual:     true,
			wantSameShape: true,
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
