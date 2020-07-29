package parser

// Parser is ...
type Parser interface {
	Parse(expr string) ([]Token, error)
}
