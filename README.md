<div align="center" id=badges>

[![Go Reference](https://pkg.go.dev/badge/github.com/MarkRosemaker/openapi-compare.svg)](https://pkg.go.dev/github.com/MarkRosemaker/openapi-compare)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarkRosemaker/openapi-compare)](https://goreportcard.com/report/github.com/MarkRosemaker/openapi-compare)
![Code Coverage](https://img.shields.io/badge/coverage-73.9%25-yellowgreen)
[![License: Apache](https://img.shields.io/badge/License-Apache-yellow.svg)](./LICENSE)

</div>

<p align="center">
  <img alt="A gopher holding up two nearly identical sheets of paper, checking them against each other" src=openapi-compare.jpg width=500>
</p>

<h3 align="center">
  Decide when two pieces of an API spec are the same.
</h3>

`openapi-compare` answers a question that sounds trivial and isn't: are these two
objects in an [OpenAPI 3.x](https://spec.openapis.org/oas/v3.1.0) specification the
same? Two schemas can describe exactly the same JSON and still differ in their
`title`, their `description`, or their `example`. Whether that counts as "the same"
depends entirely on what you are about to do with the answer.

So this module doesn't offer one comparison. It offers a small set of comparisons
with clearly stated semantics, and lets the caller pick.

## Introduction

The comparisons here are extracted from the tools that needed them, not invented up
front. [`openapi-compress`](https://github.com/MarkRosemaker/openapi-compress) uses
them to decide when two component schemas can be collapsed into one — a decision
that must ignore cosmetic differences to be useful, and must not ignore semantic
ones to be correct.

Each kind of specification object gets its own subpackage, so function names read
naturally at the call site and don't have to compete for a single flat namespace:

```go
schema.Equal(a, b)     // not compare.SchemaEqual(a, b)
```

Today that means the `schema` package. Comparisons for operations, paths, and other
objects can be added the same way as they are actually needed.

## Features

### `schema` — comparing `*openapi.Schema`

| Function | Semantics |
|---|---|
| `schema.Equal(a, b)` | Full fidelity. Every field must match, including `title`, `description`, and `default`. |
| `schema.SameShape(a, b)` | Validation equivalence. Reports whether the same JSON documents would pass or fail against both schemas. |

`SameShape` ignores documentation-only fields — `title`, `description`, `default`,
and `example` — because none of them constrain an instance. It does **not** ignore
specification extensions: a custom `x-` extension can carry meaning that a generic
comparison has no way to reason about, so schemas whose extensions differ are never
reported as the same shape.

Both functions recurse consistently. `Equal` recurses through `Equal`, `SameShape`
through `SameShape`, so a difference buried three levels deep inside a property is
surfaced by exactly the comparison that cares about it. Composition keywords
(`allOf`, `oneOf`, `anyOf`, `not`), `items`, `properties`, and
`additionalProperties` are all covered.

`example` is ignored by both. Per the OpenAPI and JSON Schema specifications it is
documentation only and never affects what an instance validates against.

## Usage

```go
import (
    "github.com/MarkRosemaker/openapi"
    "github.com/MarkRosemaker/openapi-compare/schema"
)

a := &openapi.Schema{Type: openapi.TypeString, Description: "the user's email"}
b := &openapi.Schema{Type: openapi.TypeString, Description: "email address"}

schema.Equal(a, b)     // false — the descriptions differ
schema.SameShape(a, b) // true  — the same strings validate against both
```

Use `Equal` when a difference of any kind matters, such as detecting whether a
document changed. Use `SameShape` when deciding whether two definitions are
interchangeable, such as deduplicating components.

## The openapi family

| Module | Purpose |
|---|---|
| [openapi](https://github.com/MarkRosemaker/openapi) | Parse, validate, and write OpenAPI 3.x specifications |
| **openapi-compare** (this module) | Compare specification objects — exact equality and shape equivalence |
| [openapi-edit](https://github.com/MarkRosemaker/openapi-edit) | Safe structural edits, such as renaming a schema and rewriting every `$ref` to it |
| [openapi-flatten](https://github.com/MarkRosemaker/openapi-flatten) | Promote inline definitions into named `components` entries |
| [openapi-compress](https://github.com/MarkRosemaker/openapi-compress) | Deduplicate and merge equivalent component schemas |
| [openapi-merge](https://github.com/MarkRosemaker/openapi-merge) | Merge schemas that were inferred independently from different samples |
| [openapi-enrich](https://github.com/MarkRosemaker/openapi-enrich) | Infer specification content from observed HTTP traffic |
| [openapi-codegen](https://github.com/MarkRosemaker/openapi-codegen) | Generate Go types, clients, and servers from a specification |

Note the difference from [`openapi-merge`](https://github.com/MarkRosemaker/openapi-merge):
this module only ever *reports* on two objects, it never modifies them. Merging two
schemas into a single wider one is what `openapi-merge` is for.

## Additional Information

- [**Go Reference**](https://pkg.go.dev/github.com/MarkRosemaker/openapi-compare): The Go reference documentation for the openapi-compare package.
- [**Go Report Card**](https://goreportcard.com/report/github.com/MarkRosemaker/openapi-compare): Check the code quality report.

Requires Go with `GOEXPERIMENT=jsonv2` (set via `go env -w GOEXPERIMENT=jsonv2`), inherited from [`openapi`](https://github.com/MarkRosemaker/openapi).

## Contributing

If you have any contributions to make, please submit a pull request or open an issue on the [GitHub repository](https://github.com/MarkRosemaker/openapi-compare).

## License

This project is licensed under the [Apache 2.0 License](./LICENSE).
