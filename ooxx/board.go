package game

import (
	"fmt"
	"strings"
)

type Board struct {
	board string
}

func NewBoard() *Board {
	return &Board{
		board: "_________",
	}
}

func (b *Board) String() string {
	var builder strings.Builder
	builder.WriteString("+---+---+---+\n")
	for i := 0; i < len(b.board); i++ {
		if b.board[i] == SYMBOL_EMPTY {
			builder.WriteString("|   ")
		} else {
			builder.WriteString(fmt.Sprintf("| %c ", b.board[i]))
		}
		if (i+1)%3 == 0 {
			builder.WriteString("|\n")
			builder.WriteString("+---+---+---+\n")
		}
	}

	return builder.String()
}
