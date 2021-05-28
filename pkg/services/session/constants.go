package session

// Cookie names.
const (
	PreSessionAccessTokenCookie = "__Host_Gigamono_Pre_Session_Access_Token_JWT"
	SessionAccessTokenCookie    = "__Host_Gigamono_Session_Access_Token_JWT"
	SessionRefreshTokenCookie   = "__Host_Gigamono_Session_Refresh_Token_JWT"
)

// Header names.
const (
	PreSessionCSRFHeader = "X-Pre-Session-CSRF-ID"
	SessionCSRFHeader    = "X-Session-CSRF-ID"
)
