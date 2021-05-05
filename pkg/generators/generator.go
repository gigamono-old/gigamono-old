package generators

// Generator abstracts how code is generated from a target.
type Generator interface {
	Generate(source string, opts ...interface{}) (string, error)
}
