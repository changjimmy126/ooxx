package main

import "fmt"

func main () {
        board := "OO_XO_XOX"
        printBoard(board)
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