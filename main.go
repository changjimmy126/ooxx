package main

import (
	"errors"
	"fmt"
)

func main() {
	board := "_________"

	var winner rune
	isGameOver := false
	currentPlayer := SYMBOL_O
	
	for !isGameOver {
		// todo 請玩家下棋
		var newplace int
		var err error
		var newboard string

		fmt.Println("請選擇位置1~9")
		fmt.Scan(&newplace)

		newboard, err = place(board, newplace, currentPlayer)

		if err != nil {
			fmt.Println(err)
			continue
		}
		if currentPlayer == SYMBOL_O {
			currentPlayer = SYMBOL_X
		} else {
			currentPlayer = SYMBOL_O
		}

		board = newboard
		printBoard(board)

		// 檢查遊戲是否結束
		winner, isGameOver = checkGameOver(board)
		continue
	}

	if winner == SYMBOL_EMPTY {
		fmt.Println("平手")
	} else {
		fmt.Printf("玩家 %c 獲勝\n", winner)
	}
}

const (
	SYMBOL_O     = 'O'
	SYMBOL_X     = 'X'
	SYMBOL_EMPTY = '_'
)

func printBoard(board string) {
	fmt.Println("+---+---+---+")
	for i := 0; i < len(board); i++ {
		if board[i] == SYMBOL_EMPTY {
			fmt.Printf("|   ")
		} else {
			fmt.Printf("| %c ", board[i])
		}
		if (i+1)%3 == 0 {
			fmt.Println("|")
			fmt.Println("+---+---+---+")
		}
	}
}

func place(board string, position int, symbol rune) (string, error) {
	if position < 1 || position > len(board) {
		return "", errors.New("選擇的位置不存在")
	}
	byteBoard := []byte(board)
	byteBoard[position-1] = byte(symbol)
	return string(byteBoard), nil

}

// isGameOver 返回兩個參數，代表(獲勝玩家, 遊戲是否結束)，若平手則回傳(SYMBOL_EMPTY, true)
func checkGameOver(board string) (rune, bool) {
	// 判斷垂直方向上是否有玩家獲勝
	for i := 0; i < 3; i++ {
		if board[i] == board[i+3] && board[i+3] == board[i+6] && board[i] != SYMBOL_EMPTY {
			return rune(board[i]), true
		}
	}
	// 判斷水平方向上是否有玩家獲勝
	for i := 0; i < 3; i++ {
		if board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] && board[i*3] != SYMBOL_EMPTY {
			return rune(board[i*3]), true
		}
	}
	// 判斷對角方向上是否有玩家獲勝
	if board[0] == board[4] && board[4] == board[8] && board[0] != SYMBOL_EMPTY {
		return rune(board[0]), true
	} else if board[2] == board[4] && board[4] == board[6] && board[2] != SYMBOL_EMPTY {
		return rune(board[2]), true
	}

	return SYMBOL_EMPTY, false
}
