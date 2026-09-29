### `schema` — comparing `*openapi.Schema`

| Function | Semantics |
|---|---|
| `schema.Equal(a, b)` | Full fidelity. Every field must match, including `title`, `description`, and `default`. |
| `schema.SameShape(a, b)` | Validation equivalence. Reports whether the same JSON documents would pass or fail against both schemas. |

`SameShape` ignores documentation-only fields — `title`, `description`, `default`,
`deprecated`, `example` and `examples` — because none of them constrain an instance.
It does **not** ignore specification extensions or the `discriminator`: an `x-`
extension can carry meaning that a generic comparison has no way to reason about,
and a discriminator decides how code tells the alternatives apart, so schemas that
differ in either are never reported as the same shape. A discriminator's `mapping` counts too, with a schema's name and a reference to it treated alike.

Both functions recurse consistently. `Equal` recurses through `Equal`, `SameShape`
through `SameShape`, so a difference buried three levels deep inside a property is
surfaced by exactly the comparison that cares about it. Composition keywords
(`allOf`, `oneOf`, `anyOf`, `not`), `prefixItems`, `items`, `properties`,
`additionalProperties` and `propertyNames` are all covered.

A `$ref` is compared by where it points, not by following it: two references to
the same schema match, and a schema that refers to itself compares without
recursing forever. Keywords beside a `$ref` are compared like any other.

`SameShape` treats an absent `additionalProperties`, `true`, and the empty schema
as the same, since each accepts any extra property. `Equal` tells them apart.

`example` and `examples` are ignored by both. Per the OpenAPI and JSON Schema specifications it is
documentation only and never affects what an instance validates against.
