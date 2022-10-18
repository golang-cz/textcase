# textcase

[![GoDoc Widget]][GoDoc]

Golang pkg to convert any text input to camelCase, PascalCase or snake_case naming convention. Removes all whitespaces and special characters. Supports Unicode characters.

## Usage
```go
import "github.com/golang-cz/textcase"
```

```go
textcase.CamelCase("Hello World!")
// helloWorld
```

```go
textcase.PascalCase("Hello World!")
// HelloWorld
```

```go
textcase.SnakeCase("Hello World!")
// hello_world
```

## Unicode support
```go
textcase.CamelCase("Háčky, čárky. Příliš žluťoučký kůň úpěl ďábelské ódy.")
// háčkyČárkyPřílišŽluťoučkýKůňÚpělĎábelskéÓdy
```

# License
Licensed under [MIT License](./LICENSE)

[GoDoc]: https://godoc.org/github.com/golang-cz/textcase
[GoDoc Widget]: https://godoc.org/github.com/golang-cz/textcase?status.svg
