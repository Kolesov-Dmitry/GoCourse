package chess

var knightMoves = [...][2]int { { 2, 1 }, { 2, -1 }, { -2, 1 }, { -2, -1 }, { 1, 2 }, { -1, 2 }, { 1, -2 }, { -1, -2 } }

func GetKnightMoves(space *Space) []*Space {	
	result := make([]*Space, 0, 8)

	for _, move := range knightMoves {
		space, err := space.Move(move[0], move[1])
		if err == nil {
			result = append(result, space)
		}
	}

	return result
}