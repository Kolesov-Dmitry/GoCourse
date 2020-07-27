package parser

// TokenType data type
type TokenType int

// Token types enumeration
const (
	Digit TokenType = iota
	PlusOperator
	MinusOperator
	MulOperator
	DivOperator
	OpenParentheses
	CloseParentheses
)

// Token representation
type Token struct {
	Type  TokenType
	Value string
}
