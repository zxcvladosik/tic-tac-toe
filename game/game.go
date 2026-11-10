package game

import (
	"fmt"
	"math/rand"
)

type Board struct {
	cells [3][3]rune
}

func NewBoard() *Board {
	board := &Board{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.cells[i][j] = ' '
		}
	}

	return board
}

func (b *Board) Display() {
	fmt.Println("\n    1   2   3")
	fmt.Println("  +---+---+---+")

	for i := 0; i < 3; i++ {
		fmt.Printf("%d |", i+1)

		for j := 0; j < 3; j++ {
			fmt.Printf(" %c |", b.cells[i][j])
		}

		fmt.Println()
		fmt.Println("  +---+---+---+")
	}
}

func (b *Board) MakeMove(row, col int, mark rune) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return false
	}

	if b.cells[row][col] != ' ' {
		return false
	}

	b.cells[row][col] = mark
	return true
}

func (b *Board) Winner() rune {
	for i := 0; i < 3; i++ {
		if b.cells[i][0] != ' ' &&
			b.cells[i][0] == b.cells[i][1] &&
			b.cells[i][1] == b.cells[i][2] {
			return b.cells[i][0]
		}

		if b.cells[0][i] != ' ' &&
			b.cells[0][i] == b.cells[1][i] &&
			b.cells[1][i] == b.cells[2][i] {
			return b.cells[0][i]
		}
	}

	if b.cells[0][0] != ' ' &&
		b.cells[0][0] == b.cells[1][1] &&
		b.cells[1][1] == b.cells[2][2] {
		return b.cells[0][0]
	}

	if b.cells[0][2] != ' ' &&
		b.cells[0][2] == b.cells[1][1] &&
		b.cells[1][1] == b.cells[2][0] {
		return b.cells[0][2]
	}

	return ' '
}

func (b *Board) IsDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.cells[i][j] == ' ' {
				return false
			}
		}
	}

	return b.Winner() == ' '
}

func (b *Board) ComputerMove(mark rune) bool {
	opponent := 'X'

	if mark == 'X' {
		opponent = 'O'
	}

	// 1. Проверяем, может ли компьютер победить.
	move, found := b.findWinningMove(mark)

	if found {
		return b.MakeMove(move[0], move[1], mark)
	}

	// 2. Проверяем, нужно ли блокировать противника.
	move, found = b.findWinningMove(opponent)

	if found {
		return b.MakeMove(move[0], move[1], mark)
	}

	// 3. Если угроз нет, выбираем случайную клетку.
	var available [][2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.cells[i][j] == ' ' {
				available = append(available, [2]int{i, j})
			}
		}
	}

	if len(available) == 0 {
		return false
	}

	move = available[rand.Intn(len(available))]
	return b.MakeMove(move[0], move[1], mark)
}

// findWinningMove ищет клетку,
// которая позволит указанному игроку победить.
func (b *Board) findWinningMove(mark rune) ([2]int, bool) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			// Пропускаем занятые клетки.
			if b.cells[i][j] != ' ' {
				continue
			}

			// Временно ставим символ.
			b.cells[i][j] = mark

			// Проверяем, появилась ли победа.
			winner := b.Winner()

			// Возвращаем клетку в исходное состояние.
			b.cells[i][j] = ' '

			if winner == mark {
				return [2]int{i, j}, true
			}
		}
	}

	return [2]int{}, false
}
