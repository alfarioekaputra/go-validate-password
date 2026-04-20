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

| Rule | Requirement |
|------|-------------|
| Length | Minimum 14 characters (configurable) |
| Uppercase | At least one uppercase letter (A–Z) |
| Lowercase | At least one lowercase letter (a–z) |
| Digit | At least one numeric digit (0–9) |
| Special character | At least one punctuation or symbol character (e.g. `!@#$%^&*`) |

## Return Values

Both functions return `(bool, string)`:

- `(true, "Password is valid")` — password meets all criteria
- `(false, "<reason>")` — password fails, with a human-readable explanation

## Functions

### `ValidatePassword`

Validates using the default minimum length of 14 characters.

```go
ok, msg := validate.ValidatePassword("MyP@ssw0rd!2024")
```

### `ValidatePasswordWithOptions`

Validates with a custom minimum length via `Options`.

```go
opts := validate.Options{MinLength: 8}
ok, msg := validate.ValidatePasswordWithOptions("MyP@ss1!", opts)
```

If `MinLength` is 0 or not set, it defaults to 14.

## Example

```go
package main

import (
    "fmt"
    "github.com/alfarioekaputra/go-validate-password"
)

func main() {
    // Default: minimum 14 characters
    ok, msg := validate.ValidatePassword("ValidP@ssw0rd!123")
    fmt.Printf("[%v] %s\n", ok, msg)

    // Custom: minimum 8 characters
    opts := validate.Options{MinLength: 8}
    ok, msg = validate.ValidatePasswordWithOptions("MyP@ss1!", opts)
    fmt.Printf("[%v] %s\n", ok, msg)
}
```

## License

MIT
