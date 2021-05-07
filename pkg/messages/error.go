package messages

import "fmt"

// Error maps a key to error message.
var Error = map[string]interface{}{
	"internal-system-error": "internal system error",
	"input-validation": func(n string) string {
		return fmt.Sprintf("validation error: `%v` input has a wrong value", n)
	},
	"grant-type": "unsupported grant type",
}
