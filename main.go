package main

import (
	"fmt"

	game "ooxx/ooxx"
)

func main() {
	game := game.NewGame()

	for !game.CheckGameOver() {
		var newplace int

		fmt.Println("請選擇位置1~9")
		fmt.Scan(&newplace)

		if err := game.Place(newplace, game.CurrentPlayer); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(game.Board)
	}

	fmt.Println(game.GetWinner())
}
