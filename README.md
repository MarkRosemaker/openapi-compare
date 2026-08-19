# openapi-compare

Comparisons for [github.com/MarkRosemaker/openapi](https://github.com/MarkRosemaker/openapi) objects, shared between [openapi-compress](https://github.com/MarkRosemaker/openapi-compress) and [openapi-merge](https://github.com/MarkRosemaker/openapi-merge).

Each kind of OpenAPI object gets its own subpackage, so function names don't have to stutter or compete for a single flat namespace:

- `schema` — `Equal`, `SameShape` for `*openapi.Schema`.
