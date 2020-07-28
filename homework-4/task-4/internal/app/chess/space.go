package chess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const chessBoardSize int = 8
var chessBoardColumns = [...]byte { ' ', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H' }	

// Space is a cell on a chessboard
type Space struct {
	col int
	row int
}

// NewSpace creates A1 space on a chessboard
func NewSpace() *Space {	
	return &Space { col: 1, row: 1 }
}

// SpaceFromString creates a Space instance from position presented by a string like "A1", "E5", ...
// Input:
//  - str is a string representation of the position on a chessboard
// Output:
//  - Space instance in case of success
//  - otherwise, returns error
func SpaceFromString(str string) (*Space, error) {
	if len(str) != 2 {
		return nil, errors.New(fmt.Sprintf("Invalid space '%s'", str))
	}

	str = strings.ToUpper(str)

	row := int(str[0] - 'A') + 1
	col, _ := strconv.ParseInt(string(str[1]), 10, 32)

	space := NewSpace()
	space, err := space.Move(row - 1, int(col) - 1)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid space '%s'", str))
	}

	return space, nil
}

// Move creates a new Space instance obtained by offsetting the current space
// Input:
//  - rowOffset is the value to offset the row of current space
//  - colOffset is the value to offset the column of current space
// Output:
//  - new Space instance in case of success
//  - otherwise, returns error
func (c *Space) Move(rowOffset int, colOffset int) (*Space, error) {
	row := c.row + rowOffset
	col := c.col + colOffset

	if row < 1 || row > chessBoardSize || col < 1 || col > chessBoardSize {
		return nil, errors.New("Invalid space")
	}

	return &Space { col: col, row: row }, nil
}

// String is the Stringer interface implementation for Space type
func (c Space) String() string {
	return fmt.Sprintf("%s%d", string(chessBoardColumns[c.row]), c.col)
}