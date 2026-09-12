---
tagline: Decide when two pieces of an API spec are the same.
logo:
    alt: A gopher holding up two OpenAPI schemas, checking them against each other
    source: openapi-compare.jpg
    width: 500
---

<div align="center" id=badges>

![Code Coverage](https://img.shields.io/badge/coverage-73.9%25-yellowgreen)

</div>





`openapi-compare` answers a question that sounds trivial and isn't: are these two
objects in an [OpenAPI 3.x](https://spec.openapis.org/oas/v3.1.0) specification the
same? Two schemas can describe exactly the same JSON and still differ in their
`title`, their `description`, or their `example`. Whether that counts as "the same"
depends entirely on what you are about to do with the answer.

So this module doesn't offer one comparison. It offers a small set of comparisons
with clearly stated semantics, and lets the caller pick.
