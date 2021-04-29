package compilers

// Compiler abstracts how a source is compiled to a target.
type Compiler interface {
	Compile(source string, opts ...interface{}) (string, error)
}
