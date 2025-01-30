package main

import "fmt"

func main () {
        board := "_________"

		var winner rune
        isGameOver := false

        for !isGameOver {
                // todo 請玩家下棋
				var newplace int
				
				fmt.Println("請選擇位置1~9")
				fmt.Scan(&newplace)
					if newplace < 1 || newplace > 9 {
						fmt.Println("無效輸入，請重複")
						continue
					}
			place(board , newplace , SYMBOL_O)
			board = place(board, newplace, SYMBOL_O)
				printBoard(board)
				

                // 檢查遊戲是否結束
                winner, isGameOver = checkGameOver(board)
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

func place(board string, position int, symbol rune) string {
        byteBoard := []byte(board)
        byteBoard[position-1] = byte(symbol)
        return string(byteBoard)
}

// isGameOver 返回兩個參數，代表(獲勝玩家, 遊戲是否結束)，若平手則回傳(SYMBOL_EMPTY, true)
func checkGameOver(board string) (rune, bool) {
	// todo 判斷垂直方向上是否有玩家獲勝
    for i := 0; i < 3; i++ {
		if board[i] == board[i+3] && board[i+3] == board[i+6] {
			return rune(board[i]), true
		}
	}
	// todo 判斷水平方向上是否有玩家獲勝
	for i := 0; i < 3; i++ {
		if board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] {
			return rune(board[i*3]), true
		}
	}
	// todo 判斷對角方向上是否有玩家獲勝
		if board[0] == board[4] && board[4] == board[8] {
			return rune(board[0]), true
		} else if board[2] == board[4] && board[4] == board[6] {
			return rune(board[2]), true
		}

	return SYMBOL_EMPTY, false
}