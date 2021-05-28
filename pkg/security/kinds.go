package security

// ActionKind represents a Claim Action kind
type ActionKind string

// ...
const (
	PreSession     ActionKind = "PreSession"
	SessionAccess  ActionKind = "SessionAccess"
	SessionRefresh ActionKind = "SessionRefresh"
)
