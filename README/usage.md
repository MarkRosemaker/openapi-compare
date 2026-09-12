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
