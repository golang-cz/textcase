# textcase

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis]

Convert any text input to ASCII CamelCase or snake_case via a fast state machine parser.

## Usage:
```go
import "github.com/golang-cz/textcase"
```

```go
textcase.CamelCase("Hello World!!")
// HelloWorld
```

```go
textcase.SnakeCase("Hello World!!")
// hello_world
```

# License
Licensed under [MIT License](./LICENSE)

[GoDoc]: https://godoc.org/github.com/golang-cz/textcase
[GoDoc Widget]: https://godoc.org/github.com/golang-cz/textcase?status.svg
[Travis]: https://travis-ci.org/golang-cz/textcase
[Travis Widget]: https://travis-ci.org/golang-cz/textcase.svg?branch=master
