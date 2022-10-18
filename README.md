# textcase

[![GoDoc Widget]][GoDoc]

Golang pkg to convert any text input to **camelCase**, **PascalCase**, **snake_case** or **kebab-case** naming convention. Removes all whitespaces and special characters. Supports Unicode characters.

## Usage
```go
import "github.com/golang-cz/textcase"

textcase.CamelCase("Hello World!")
// helloWorld

textcase.PascalCase("Hello World!")
// HelloWorld

textcase.SnakeCase("Hello World!")
// hello_world

textcase.KebabCase("Hello World!")
// hello-world
```

## Unicode support
```go
textcase.CamelCase("Háčky, čárky. Příliš žluťoučký kůň úpěl ďábelské ódy.")
// háčkyČárkyPřílišŽluťoučkýKůňÚpělĎábelskéÓdy
```

### Possible Unicode limitations
```go
textcase.CamelCase("Here comes O'Brian")
// hereComesOBrian
```

This package doesn't implement language-specific case mappers, such as [golang.org/x/text/cases](https://pkg.go.dev/golang.org/x/text/cases), and thus comes with a similar limitation to [strings.Title()](https://pkg.go.dev/strings#Title). But given the likely use cases of this package, we deliberately chose English version `hereComesOBrian` over `hereComesObrian` for the above `Here comes O'Brian` input.

# License
Licensed under [MIT License](./LICENSE)

[GoDoc]: https://godoc.org/github.com/golang-cz/textcase
[GoDoc Widget]: https://godoc.org/github.com/golang-cz/textcase?status.svg
