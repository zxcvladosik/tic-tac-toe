package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/tic-tac-toe/game"
	"example.com/tic-tac-toe/statistics"
)

const statsFile = "stats.json"

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	stats, err := statistics.Load(statsFile)

	if err != nil {
		fmt.Println("Ошибка загрузки статистики:", err)
		stats = statistics.Stats{}
	}

	fmt.Println("================================")
	fmt.Println("       КРЕСТИКИ-НОЛИКИ")
	fmt.Println("================================")

	for {
		fmt.Println("\nГлавное меню:")
		fmt.Println("1. Играть против другого игрока")
		fmt.Println("2. Играть против компьютера")
		fmt.Println("3. Посмотреть статистику")
		fmt.Println("0. Выход")

		choice := readNumber("Выберите действие: ")

		switch choice {
		case 1:
			playGame("players", &stats)
		case 2:
			playGame("computer", &stats)
		case 3:
			showStatistics(stats)
		case 0:
			fmt.Println("Спасибо за игру!")
			return
		default:
			fmt.Println("Некорректный выбор.")
		}
	}
}

func readNumber(prompt string) int {
	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			return 0
		}

		number, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if err == nil {
			return number
		}

		fmt.Println("Введите целое число.")
	}
}

func readMove() (int, int, bool) {
	fmt.Print("Введите строку и столбец (например, 1 2): ")

	if !scanner.Scan() {
		return 0, 0, false
	}

	parts := strings.Fields(scanner.Text())

	if len(parts) != 2 {
		fmt.Println("Нужно ввести два числа.")
		return 0, 0, true
	}

	row, err1 := strconv.Atoi(parts[0])
	col, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Координаты должны быть числами.")
		return 0, 0, true
	}

	return row - 1, col - 1, true
}

func playGame(mode string, stats *statistics.Stats) {
	board := game.NewBoard()
	currentPlayer := 'X'

	fmt.Println("\nИгра началась!")

	for {
		board.Display()

		if mode == "computer" && currentPlayer == 'O' {
			fmt.Println("Ход компьютера...")

			board.ComputerMove('O')
		} else {
			fmt.Printf("Ход игрока %c\n", currentPlayer)

			for {
				row, col, ok := readMove()

				if !ok {
					fmt.Println("Игра прервана.")
					return
				}

				if board.MakeMove(row, col, currentPlayer) {
					break
				}

				fmt.Println("Эта клетка занята или координаты неверны.")
			}
		}

		winner := board.Winner()

		if winner != ' ' {
			board.Display()
			fmt.Printf("Победил игрок %c!\n", winner)

			stats.Record(string(winner), mode)
			saveStatistics(*stats)
			return
		}

		if board.IsDraw() {
			board.Display()
			fmt.Println("Ничья!")

			stats.Record("draw", mode)
			saveStatistics(*stats)
			return
		}

		if currentPlayer == 'X' {
			currentPlayer = 'O'
		} else {
			currentPlayer = 'X'
		}
	}
}

func saveStatistics(stats statistics.Stats) {
	err := statistics.Save(statsFile, stats)

	if err != nil {
		fmt.Println("Ошибка сохранения статистики:", err)
	} else {
		fmt.Println("Статистика сохранена.")
	}
}

func showStatistics(stats statistics.Stats) {
	fmt.Println("\n========== СТАТИСТИКА ==========")
	fmt.Println("Всего игр:", stats.TotalGames)
	fmt.Println("Победы первого игрока (X):", stats.Player1Wins)
	fmt.Println("Победы второго игрока (O):", stats.Player2Wins)
	fmt.Println("Победы компьютера:", stats.ComputerWins)
	fmt.Println("Ничьи:", stats.Draws)
	fmt.Println("================================")
}
