# camelcase

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
[MIT license](./LICENSE)
