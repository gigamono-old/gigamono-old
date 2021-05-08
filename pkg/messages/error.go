package messages

import "fmt"

// Error maps a key to error message.
var Error = map[string]interface{}{
	"internal": "internal system error",
	"validation": func(n string) string {
		return fmt.Sprintf("validation error: `%v` input has a wrong value", n)
	},
	"grant-type":      "unsupported grant type",
	"user-auth":       "cannot get current user session",
	"workflow-config": "invalid workflow config",
	"workflow-id":     "invalid workflow id",
}
