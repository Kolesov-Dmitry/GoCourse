package expression

import (
	"errors"
	"strconv"

	"homework.com/v4/task-3/internal/app/parser"
)

// Solver evalueates algebraic expressions
type Solver struct {
	tokens []parser.Token
	stack  []float64
}

// NewSolver makes new Solver object instance
func NewSolver() *Solver {
	return &Solver{
		tokens: nil,
		stack:  make([]float64, 0),
	}
}

// Solve evalueates algebraic expressions
// Input:
//  - expr is algebraic expressions
// Output:
//  - result of evaluation in case of success
//  - otherwise, returns error
func (s *Solver) Solve(expr string) (float64, error) {
	var p parser.ExprParser
	tokens, err := p.Parse(expr)
	if err != nil {
		return 0.0, err
	}

	var c Converter
	rpnTokens, err := c.Convert(tokens)
	if err != nil {
		return 0.0, err
	}

	s.tokens = rpnTokens
	return s.solveExpression()
}

// solveOperator performs appropriate algebraic operator using valuse from stack
// Input:
//  - token wich containt operator type
// return results if there is not enaught values on the stack
func (s *Solver) solveOperator(tokenType parser.TokenType) error {
	if len(s.stack) < 2 {
		return errors.New("Invalid expression")
	}

	a := s.stack[len(s.stack)-1]
	b := s.stack[len(s.stack)-2]
	s.stack = s.stack[:len(s.stack)-2]

	if tokenType == parser.PlusOperator {
		a += b

	} else if tokenType == parser.MinusOperator {
		a = b - a

	} else if tokenType == parser.MulOperator {
		a *= b

	} else if tokenType == parser.DivOperator {
		a = b / a

	}
	s.stack = append(s.stack, a)

	return nil
}

// solveExpression solve expression in Reversed Polish Notation
// Output:
//  - result of the expression
//  - error, if the expression is invalid
func (s *Solver) solveExpression() (float64, error) {
	for _, token := range s.tokens {
		switch token.Type {
		case parser.Digit:
			value, _ := strconv.ParseFloat(token.Value, 32)
			s.stack = append(s.stack, value)

		case parser.PlusOperator, parser.MinusOperator, parser.MulOperator, parser.DivOperator:
			if err := s.solveOperator(token.Type); err != nil {
				return 0.0, err
			}

		default:
			return 0.0, errors.New("Invalid expression")
		}
	}

	if len(s.stack) != 1 {
		return 0.0, errors.New("Invalid expression")
	}

	return s.stack[0], nil
}
