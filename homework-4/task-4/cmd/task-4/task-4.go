package main

import (	
	"fmt"	
	"homework.com/v4/task-4/internal/app/chess"
	"log"	
)

func main() {	
	fmt.Print("Введите позицию коня (C5, A3, ...): ")
	var pos string	
	fmt.Scan(&pos)
	
	knightSpace, err := chess.SpaceFromString(pos)
	if err != nil {
		log.Fatal(err)
	}

	knightMoves := chess.GetKnightMoves(knightSpace)	
	fmt.Println("Возможные ходы конём:", knightMoves)
}