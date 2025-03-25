package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// 定義遊戲物件
type Game struct {
	GameId        string `json:"game_id"`        // 局號
	Status        string `json:"status"`         // 遊戲狀態
	Board         string `json:"board"`          // 盤面
	CurrentPlayer rune   `json:"current_player"` // 當前輪到哪位玩家
	Winner        rune   `json:"winner"`         // 獲勝玩家
}

func NewGame() *Game {
	return &Game{
		GameId:        uuid.New().String(),
		Status:        STATE_IN_PROGRESS,
		Board:         INIT_BOARD,
		CurrentPlayer: SYMBOL_O,
		Winner:        SYMBOL_EMPTY,
	}
}

func (g *Game) GetWinner() string {
	if g.Winner == SYMBOL_EMPTY {
		return "平手"
	}

	return fmt.Sprintf("玩家 %c 獲勝", g.Winner)
}

func (g *Game) Place(position int, symbol rune) error {
	if position < 1 || position > len(g.Board) {
		return errors.New("選擇的位置不存在")
	}

	if g.Board[position-1] != SYMBOL_EMPTY {
		return errors.New("該位置已被占用，請重新選擇!")
	}

	if symbol != g.CurrentPlayer {
		return errors.New("輪到玩家 " + string(g.CurrentPlayer) + " 下棋")
	}

	byteBoard := []byte(g.Board)
	byteBoard[position-1] = byte(symbol)

	g.Board = string(byteBoard)

	if g.CurrentPlayer == SYMBOL_O {
		g.CurrentPlayer = SYMBOL_X
	} else {
		g.CurrentPlayer = SYMBOL_O
	}

	return nil
}

func (g *Game) CheckGameOver() bool {
	// 判斷垂直方向上是否有玩家獲勝
	for i := 0; i < 3; i++ {
		if g.Board[i] == g.Board[i+3] && g.Board[i+3] == g.Board[i+6] && g.Board[i] != SYMBOL_EMPTY {
			g.Winner = rune(g.Board[i])
			return true
		}
	}

	// 判斷水平方向上是否有玩家獲勝
	for i := 0; i < 3; i++ {
		if g.Board[i*3] == g.Board[i*3+1] && g.Board[i*3+1] == g.Board[i*3+2] && g.Board[i*3] != SYMBOL_EMPTY {
			g.Winner = rune(g.Board[i*3])
			return true
		}
	}

	// 判斷對角方向上是否有玩家獲勝
	if g.Board[0] == g.Board[4] && g.Board[4] == g.Board[8] && g.Board[0] != SYMBOL_EMPTY {
		g.Winner = rune(g.Board[0])
		return true
	} else if g.Board[2] == g.Board[4] && g.Board[4] == g.Board[6] && g.Board[2] != SYMBOL_EMPTY {
		g.Winner = rune(g.Board[2])
		return true
	}

	//判斷棋盤是否已滿則平手
	for i := 0; i < len(g.Board); i++ {
		if g.Board[i] == SYMBOL_EMPTY {
			return false
		}
	}

	g.Winner = SYMBOL_EMPTY
	return true
}

func (g *Game) DisplayBoard() string {
	var builder strings.Builder
	builder.WriteString("+---+---+---+\n")
	for i := 0; i < len(g.Board); i++ {
		if g.Board[i] == SYMBOL_EMPTY {
			builder.WriteString("|   ")
		} else {
			builder.WriteString(fmt.Sprintf("| %c ", g.Board[i]))
		}
		if (i+1)%3 == 0 {
			builder.WriteString("|\n")
			builder.WriteString("+---+---+---+\n")
		}
	}

	return builder.String()
}
