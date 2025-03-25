package main

import (
	"fmt"

	game "ooxx/ooxx"

	"github.com/gin-gonic/gin"
)

// 使用 map 來儲存遊戲物件（key 為遊戲 ID）
var games = make(map[string]*game.Game)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/games", createGame)
		api.GET("/games/:id", getGame)
		api.POST("/games/:id/moves/:position", move)
	}

	fmt.Println("伺服器運行於 :8080")
	r.Run(":8080")
}

// createGame 實作創建遊戲邏輯
func createGame(c *gin.Context) {
	// 創建遊戲物件
	game := game.NewGame()

	// 將遊戲物件存入 map
	games[game.GameId] = game

	// 回傳遊戲物件
	c.JSON(200, game)
}

// getGame 實作取得遊戲狀態邏輯
func getGame(c *gin.Context) {
	// 取得遊戲 ID
	gameId := c.Param("id")

	// 從 map 中取得遊戲物件
	game, ok := games[gameId]

	// 如果遊戲物件不存在，回傳 404
	if !ok {
		c.JSON(404, gin.H{
			"message": "遊戲不存在",
		})
		return
	}

	// 回傳遊戲物件
	c.JSON(200, game)
}

// move 實作落子邏輯
func move(c *gin.Context) {
	// 取得遊戲 ID
	gameId := c.Param("id")

	// 從 map 中取得遊戲物件
	game, ok := games[gameId]

	// 如果遊戲物件不存在，回傳 404
	if !ok {
		c.JSON(404, gin.H{
			"message": "遊戲不存在",
		})
		return
	}

	// todo 取得玩家落子位置

	// todo 更新遊戲物件（修改盤面）

	// 回傳遊戲物件
	c.JSON(200, game)
}

// func main() {
// 	game := game.NewGame()

// 	for !game.CheckGameOver() {
// 		var newplace int

// 		fmt.Println("請選擇位置1~9")
// 		fmt.Scan(&newplace)

// 		if err := game.Place(newplace, game.CurrentPlayer); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		fmt.Println(game.Board)
// 	}

// 	fmt.Println(game.GetWinner())
// }
