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
