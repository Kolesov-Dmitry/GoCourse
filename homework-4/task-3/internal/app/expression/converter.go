package expression

import (
	"errors"

	"homework.com/v4/task-3/internal/app/parser"
)

// operations priorities
var priorities = map[parser.TokenType]int{
	parser.PlusOperator:    1,
	parser.MinusOperator:   1,
	parser.MulOperator:     2,
	parser.DivOperator:     2,
	parser.OpenParentheses: 0,
}

// Converter structure
type Converter struct {
	output []parser.Token
	stack  []parser.Token
}

// Convert expression from straight algebraic notation to Reversed Polish Notation
// See wiki for details
// Input:
//  - straight algebraic notation expression tokens
// Output:
//  - tokens in Reversed Polish Notation
//  - error, if convertion is failed
func (c *Converter) Convert(tokens []parser.Token) ([]parser.Token, error) {
	c.output = make([]parser.Token, 0)
	c.stack = make([]parser.Token, 0)

	for _, token := range tokens {
		if token.Type == parser.Digit {
			c.output = append(c.output, token)

		} else if token.Type != parser.OpenParentheses && token.Type != parser.CloseParentheses {
			c.popUntil(priorities[token.Type])
			c.stack = append(c.stack, token)

		} else if token.Type == parser.OpenParentheses {
			c.stack = append(c.stack, token)

		} else if token.Type == parser.CloseParentheses {
			if len(c.stack) == 0 {
				return nil, errors.New("Invalid expression")
			}

			for c.stack[len(c.stack)-1].Type != parser.OpenParentheses {
				c.output = append(c.output, c.stack[len(c.stack)-1])
				c.stack = c.stack[:len(c.stack)-1]
				if len(c.stack) == 0 {
					return nil, errors.New("Invalid expression")
				}
			}

			c.stack = c.stack[:len(c.stack)-1]
		}
	}
	c.popUntil(0)

	return c.output, nil
}

// popUntil pops token from stack to output utill token with 'priority' priority
// Input:
//  - token priority
func (c *Converter) popUntil(priority int) {
	for len(c.stack) > 0 && priorities[c.stack[len(c.stack)-1].Type] >= priority {
		c.output = append(c.output, c.stack[len(c.stack)-1])
		c.stack = c.stack[:len(c.stack)-1]
	}
}
