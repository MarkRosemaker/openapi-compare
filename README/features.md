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
