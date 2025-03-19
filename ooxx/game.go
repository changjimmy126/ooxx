package game

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// 定義遊戲物件
type Game struct {
	GameId        string // 局號
	Status        string // 遊戲狀態
	Board         *Board // 盤面
	CurrentPlayer rune   // 當前輪到哪位玩家
	Winner        rune   // 獲勝玩家
}

func NewGame() *Game {
	return &Game{
		GameId:        uuid.New().String(),
		Status:        STATE_IN_PROGRESS,
		Board:         NewBoard(),
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
	if position < 1 || position > len(g.Board.board) {
		return errors.New("選擇的位置不存在")
	}

	if g.Board.board[position-1] != SYMBOL_EMPTY {
		return errors.New("該位置已被占用，請重新選擇!")
	}

	if symbol != g.CurrentPlayer {
		return errors.New("輪到玩家 " + string(g.CurrentPlayer) + " 下棋")
	}

	byteBoard := []byte(g.Board.board)
	byteBoard[position-1] = byte(symbol)

	g.Board.board = string(byteBoard)

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
		if g.Board.board[i] == g.Board.board[i+3] && g.Board.board[i+3] == g.Board.board[i+6] && g.Board.board[i] != SYMBOL_EMPTY {
			g.Winner = rune(g.Board.board[i])
			return true
		}
	}

	// 判斷水平方向上是否有玩家獲勝
	for i := 0; i < 3; i++ {
		if g.Board.board[i*3] == g.Board.board[i*3+1] && g.Board.board[i*3+1] == g.Board.board[i*3+2] && g.Board.board[i*3] != SYMBOL_EMPTY {
			g.Winner = rune(g.Board.board[i*3])
			return true
		}
	}

	// 判斷對角方向上是否有玩家獲勝
	if g.Board.board[0] == g.Board.board[4] && g.Board.board[4] == g.Board.board[8] && g.Board.board[0] != SYMBOL_EMPTY {
		g.Winner = rune(g.Board.board[0])
		return true
	} else if g.Board.board[2] == g.Board.board[4] && g.Board.board[4] == g.Board.board[6] && g.Board.board[2] != SYMBOL_EMPTY {
		g.Winner = rune(g.Board.board[2])
		return true
	}

	//判斷棋盤是否已滿則平手
	for i := 0; i < len(g.Board.board); i++ {
		if g.Board.board[i] == SYMBOL_EMPTY {
			return false
		}
	}

	g.Winner = SYMBOL_EMPTY
	return true
}
