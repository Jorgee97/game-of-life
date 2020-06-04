package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)
import "time"

// DeadState ...
func DeadState(width, height int) (a [][]int) {
	a = make([][]int, height)
	for i := range a {
		a[i] = make([]int, width)
	}
	return
}

// TODO: How the fuck do i test this D:
func Randomize(percent int) int {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(100)

	if rn > percent {
		return 1
	}
	return 0
}

// TODO: How the fuck do i test this D:
func RandomState(width, height, percent int) (a [][]int) {
	board := DeadState(width, height)

	for row := range board {
		for col := range board[row] {
			board[row][col] = Randomize(percent)
		}
	}
	return board
}

func Render(board [][]int) {
	clearConsole()
	for row := range board {
		fmt.Printf("|")
		for col := range board[row] {
			if board[row][col] == 0 {
				fmt.Printf(" ")
				continue
			}
			fmt.Printf("#")
		}
		fmt.Printf("|\n")
	}
}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	_ = c.Run()
}

func NextBoardState(board [][]int) (updatedBoard [][]int) {
	updatedBoard = DeadState(len(board[0]), len(board))
	for i := range board {
		for j := range board[i] {
			counter := 0

			// Validate diagonal top left
			if i-1 >= 0 && j-1 >= 0 {
				if board[i-1][j-1] == 1 {
					counter++
				}
			}

			// Validate top
			if i-1 >= 0 {
				if board[i-1][j] == 1 {
					counter++
				}
			}

			// Validate diagonal top right
			if i-1 >= 0 && j+1 < len(board[i]) {
				if board[i-1][j+1] == 1 {
					counter++
				}
			}

			// Validate left center
			if j-1 >= 0 {
				if board[i][j-1] == 1 {
					counter++
				}
			}

			// Validate right center
			if j+1 < len(board[i]) {
				if board[i][j+1] == 1 {
					counter++
				}
			}

			// validate diagonal bottom left
			if i+1 < len(board) && j-1 >= 0 {
				if board[i+1][j-1] == 1 {
					counter++
				}
			}

			// validate diagonal bottom right
			if i+1 < len(board) && j+1 < len(board[i]) {
				if board[i+1][j+1] == 1 {
					counter++
				}
			}

			// validate bottom
			if i+1 < len(board) {
				if board[i+1][j] == 1 {
					counter++
				}
			}

			// life or dead validation
			if counter == 0 || counter == 1 {
				updatedBoard[i][j] = 0
			} else if counter == 2 && board[i][j] == 1 {
				updatedBoard[i][j] = 1
			} else if counter == 3 {
				updatedBoard[i][j] = 1
			} else if counter > 3 {
				updatedBoard[i][j] = 0
			}
		}
	}

	return updatedBoard
}

func main() {
	board := RandomState(150, 50, 35)

	for {
		time.Sleep(650 * time.Millisecond)
		board = NextBoardState(board)
		Render(board)
	}
}
