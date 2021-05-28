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
	"session-access-token-cookie":     "session access token invalid",
	"pre-session-csrf-tamper":         "pre-session csrf ids or signature does not match",
	"session-csrf-tamper":             "session csrf ids or signature does not match",
	"pre-session-csrf-expired":        "pre-session csrf id expired",
	"session-csrf-expired":            "session csrf id expired",
	"pre-session-access-token-tamper": "pre-session access token has the wrong action type",
	"session-access-token-tamper":     "session access token has the wrong action type",
	"pre-session":                     "cannot authenticate pre-session user",
	"authenticate-user":               "cannot authenticate user",
}
