package messages

import "fmt"

// Error maps a key to error message.
var Error = map[string]interface{}{
	"validation": func(n string) string {
		return fmt.Sprintf("validation error: `%v` input has a wrong value", n)
	},
	"internal":                        "internal system error",
	"user-auth":                       "cannot get current user session",
	"basic-auth":                      "invalid basic authentication request",
	"signup":                          "cannot create user",
	"signin":                          "cannot sign user in",
	"workflow-config":                 "invalid workflow config",
	"workflow-id":                     "invalid workflow id",
	"pre-session-access-token-cookie": "pre-session access token invalid",
	"pre-session-csrf-tamper":        "pre-session csrf ids or signature does not match",
	"pre-session-csrf-expired":        "pre-session csrf id expired",
}
