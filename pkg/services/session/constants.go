package session

// Cookie names.
const (
	PreSessionAccessTokenCookie = "__Host_Gigamono_Pre_Session_Access_Token_JWT"
	AccessTokenCookie           = "__Host_Gigamono_Access_Token_JWT"
	RefreshTokenCookie          = "__Host_Gigamono_Refresh_Token_JWT"
)

// Header names.
const (
	PreSessionCSRFHeader = "X-Pre-Session-CSRF-ID"
	SessionCSRFHeader    = "X-Session-CSRF-ID"
)
