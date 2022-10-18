# textcase

[![GoDoc Widget]][GoDoc]

Golang pkg to convert any text input to camelCase, PascalCase or snake_case via a fast state machine parser. Removes all whitespaces and special characters. Supports Unicode characters.

## Usage
```go
import "github.com/golang-cz/textcase"
```

```go
textcase.CamelCase("Hello World!")
// helloWorld
```

```go
textcase.CamelCase("Hello World!")
// HelloWorld
```

```go
textcase.SnakeCase("Hello World!")
// hello_world
```

# License
Licensed under [MIT License](./LICENSE)

[GoDoc]: https://godoc.org/github.com/golang-cz/textcase
[GoDoc Widget]: https://godoc.org/github.com/golang-cz/textcase?status.svg
