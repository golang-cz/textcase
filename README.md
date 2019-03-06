# camelcase

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis]

Golang pkg to convert any input text to CamelCase or snake_case via a fast state machine parser.

## Usage:
```go
import "github.com/VojtechVitek/camelcase"
```

```go
camelsnake.CamelCase("Hey, this TEXT will have to obey some rules!!") // HeyThisTextWillHaveToObeySomeRules
```

```go
camelsnake.SnakeCase("Hey, this TEXT will have to obey some rules!!") // hey_this_text_will_have_to_obey_some_rules
```

# License
Licensed under [MIT License](./LICENSE)

[GoDoc]: https://godoc.org/github.com/VojtechVitek/camelsnake
[GoDoc Widget]: https://godoc.org/github.com/VojtechVitek/camelsnake?status.svg
[Travis]: https://travis-ci.org/VojtechVitek/camelsnake
[Travis Widget]: https://travis-ci.org/VojtechVitek/camelsnake.svg?branch=master
