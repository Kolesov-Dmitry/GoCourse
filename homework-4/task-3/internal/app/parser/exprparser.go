package parser

import (
	"fmt"
	"unicode"
)

// list of supported operations
var operators = map[rune]TokenType{
	'-': MinusOperator,
	'+': PlusOperator,
	'*': MinusOperator,
	'/': PlusOperator,
}

// ExprParser structure
type ExprParser struct {
	buffer     string
	tokens     []Token
	dotOccured bool
}

// Parse algebraic expression into list of tokens
// Input:
//  - algebraic expression
// Output:
//  - list of tokens
//  - error, if parsing was failed
func (ep *ExprParser) Parse(expr string) ([]Token, error) {
	ep.tokens = make([]Token, 0)
	ep.buffer = ""
	ep.dotOccured = false

	for _, ch := range expr {
		if ch == '-' {
			ep.parseMinusOperator()

		} else if ch == '+' || ch == '*' || ch == '/' {
			ep.parseOperator(ch)

		} else if ch == '(' || ch == ')' {
			ep.parseParentheses(ch)

		} else if ch == ' ' {
			ep.parseBuffer()

		} else if unicode.IsDigit(ch) || ch == '.' {
			if err := ep.parseDigit(ch); err != nil {
				return nil, err
			}
		}
	}

	ep.parseBuffer()

	return ep.tokens, nil
}

// parseBuffer checks if buffer contains digit or minus operator and appends
// appropriate Token into the tokens list
func (ep *ExprParser) parseBuffer() {
	if len(ep.buffer) > 0 {
		tokenType := Digit

		if ep.buffer == "-" {
			tokenType = MinusOperator
		}

		ep.tokens = append(ep.tokens, Token{tokenType, ep.buffer})
		ep.buffer = ""
		ep.dotOccured = false
	}
}

// parseParentheses parses the buffer and appends parantheses Token into the token list
// Input:
//  - character which contains open or close parantheses symbol
func (ep *ExprParser) parseParentheses(ch rune) {
	ep.parseBuffer()

	tokenType := OpenParentheses
	if ch == ')' {
		tokenType = CloseParentheses
	}

	ep.tokens = append(ep.tokens, Token{tokenType, ""})
}

// parseOperator parses the buffer and appends operator Token into the token list
// Input:
//  - character which contains operator symbol
func (ep *ExprParser) parseOperator(ch rune) {
	ep.parseBuffer()

	tokenType := PlusOperator
	if ch == '*' {
		tokenType = MulOperator
	} else if ch == '/' {
		tokenType = DivOperator
	}

	ep.tokens = append(ep.tokens, Token{tokenType, ""})
}

// parseMinusOperator appends '-' symbol into the buffer if it's empty
// otherwise, parse the buffer and append minus operator Token into the token list
func (ep *ExprParser) parseMinusOperator() {
	if len(ep.buffer) == 0 {
		ep.buffer += "-"
	} else {
		ep.parseBuffer()
		ep.tokens = append(ep.tokens, Token{MinusOperator, ""})
	}
}

// parseDigit appends ch to the buffer if ch is digit or dot
// Input:
//  - current character from the expression
// return error if 'ch' is unexpected character
func (ep *ExprParser) parseDigit(ch rune) error {
	if unicode.IsDigit(ch) {
		ep.buffer += string(ch)
	} else if ch == '.' && len(ep.buffer) != 0 && ep.dotOccured == false {
		ep.buffer += "."
		ep.dotOccured = true
	} else {
		return fmt.Errorf("Unexpected '%s' character", string(ch))
	}

	return nil
}
