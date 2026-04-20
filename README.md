# go-validate-password

A Go library for validating passwords against strict security criteria.

## Installation

```bash
go get github.com/alfarioekaputra/go-validate-password
```

## Requirements

- Go 1.25.3+

## Usage

```go
import "github.com/alfarioekaputra/go-validate-password"

valid, message := validate.ValidatePassword("MyP@ssw0rd!2024")
if !valid {
    fmt.Println("Invalid:", message)
} else {
    fmt.Println(message) // Password is valid
}
```

## Validation Rules

A password is considered valid if it meets **all** of the following criteria:

| Rule | Requirement | Default |
|------|-------------|---------|
| Min length | Minimum number of Unicode characters | 14 |
| Max length | Maximum number of Unicode characters | 128 |
| Uppercase | At least one uppercase letter (A–Z) | — |
| Lowercase | At least one lowercase letter (a–z) | — |
| Digit | At least one decimal digit (0–9) | — |
| Special character | At least one punctuation or symbol (e.g. `!@#$%^&*`) | — |

> Length is counted in **Unicode characters (runes)**, not bytes — so multi-byte characters like `é`, `ñ`, `ü` each count as one character.

## Options

| Field | Type | Default | Notes |
|-------|------|---------|-------|
| `MinLength` | `int` | `14` | Minimum safe floor is `8` — values below are clamped |
| `MaxLength` | `int` | `128` | Set to 0 to use default |

## Return Values

Both functions return `(bool, string)`:

- `(true, "Password is valid")` — password meets all criteria
- `(false, "<reason>")` — password fails, with a human-readable explanation

## Functions

### `ValidatePassword`

Validates using default settings (min 14, max 128 characters).

```go
ok, msg := validate.ValidatePassword("MyP@ssw0rd!2024")
```

### `ValidatePasswordWithOptions`

Validates with custom min/max length via `Options`.

```go
opts := validate.Options{
    MinLength: 8,
    MaxLength: 64,
}
ok, msg := validate.ValidatePasswordWithOptions("MyP@ss1!", opts)
```

## Example

```go
package main

import (
    "fmt"
    validate "github.com/alfarioekaputra/go-validate-password"
)

func main() {
    // Default: min 14, max 128 characters
    ok, msg := validate.ValidatePassword("ValidP@ssw0rd!123")
    fmt.Printf("[%v] %s\n", ok, msg)

    // Custom: min 8, max 32 characters
    opts := validate.Options{MinLength: 8, MaxLength: 32}
    ok, msg = validate.ValidatePasswordWithOptions("MyP@ss1!", opts)
    fmt.Printf("[%v] %s\n", ok, msg)
}
```

## License

MIT
