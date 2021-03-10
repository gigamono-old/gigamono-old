package validations

// Error represents a validation error message.
type Error struct {
	FieldPath  string
	FieldName  string
	FieldValue interface{}
	Rules      string
}
