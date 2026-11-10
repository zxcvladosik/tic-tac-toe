package game

import "testing"

// Проверяем, что компьютер блокирует победу игрока.
func TestComputerBlocksPlayer(t *testing.T) {
	board := NewBoard()

	board.MakeMove(0, 0, 'X')
	board.MakeMove(0, 1, 'X')

	if !board.ComputerMove('O') {
		t.Fatal("Компьютер не сделал ход")
	}

	if board.cells[0][2] != 'O' {
		t.Fatal("Компьютер не заблокировал победу игрока")
	}
}

// Проверяем, что компьютер использует возможность победить.
func TestComputerWins(t *testing.T) {
	board := NewBoard()

	board.MakeMove(0, 0, 'O')
	board.MakeMove(0, 1, 'O')
	board.MakeMove(1, 0, 'X')

	if !board.ComputerMove('O') {
		t.Fatal("Компьютер не сделал ход")
	}

	if board.Winner() != 'O' {
		t.Fatal("Компьютер не использовал возможность победить")
	}
}
