package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Corrected Structs ---

type Game struct {
	mineCount int // Renamed from 'mine'
	board     Board
	gameState int // We can use this later (e.g., 0=playing, 1=won, 2=lost)
}

type Board struct {
	cells [][]Cell // Renamed from 'mines'
	rows  int
	cols  int
}

type Cell struct {
	isMine        bool // The missing field!
	visible       bool
	marked        bool
	neighborCount int
}

// --- Main Function (Entry Point) ---

func main() {
	// Seed the random number generator
	// We need to use a proper source, not the deprecated global one.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	rows, cols := 10, 10
	mineCount := (rows * cols) / 3 // 1/3 of the cells, as we discussed

	// Initialize the game
	game := Game{
		mineCount: mineCount,
		board:     Board{cells: make([][]Cell, rows), rows: rows, cols: cols},
		gameState: 1, // 1 = playing
	}

	// Initialize the inner slices for the board
	for i := range game.board.cells {
		game.board.cells[i] = make([]Cell, cols)
	}

	// Call the new functions
	placeMines(&game, r)
	calculateNumbers(&game)

	// --- Print the results (for testing) ---
	println("Game initialized with", game.mineCount, "mines.")
	println("Board size:", game.board.rows, "x", game.board.cols)

	println("\n--- Solution Board (for debugging) ---")
	printBoard(&game, true) // Print with all cells revealed
}

// --- Implemented Helper Functions ---

func placeMines(game *Game, r *rand.Rand) {
	placed := 0
	for placed < game.mineCount {
		row := r.Intn(game.board.rows)
		col := r.Intn(game.board.cols)

		// Only place a mine if one isn't already there
		if !game.board.cells[row][col].isMine {
			game.board.cells[row][col].isMine = true
			placed++
		}
	}
}

func calculateNumbers(game *Game) {
	for r := 0; r < game.board.rows; r++ {
		for c := 0; c < game.board.cols; c++ {
			// No need to count neighbors for a mine
			if game.board.cells[r][c].isMine {
				continue
			}

			count := 0
			// Check all 8 neighbors
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue // Skip self
					}

					nr, nc := r+dr, c+dc

					// Check if the neighbor is valid and is a mine
					if isValid(&game.board, nr, nc) && game.board.cells[nr][nc].isMine {
						count++
					}
				}
			}
			game.board.cells[r][c].neighborCount = count
		}
	}
}

// Helper to check if a coordinate is within the board
func isValid(board *Board, r, c int) bool {
	return r >= 0 && r < board.rows && c >= 0 && c < board.cols
}

// Helper to print the board
func printBoard(game *Game, revealAll bool) {
	for r := 0; r < game.board.rows; r++ {
		for c := 0; c < game.board.cols; c++ {
			cell := game.board.cells[r][c]
			if !cell.visible && !revealAll {
				fmt.Print("[ ]") // Hidden
			} else if cell.isMine {
				fmt.Print(" * ") // Mine
			} else {
				fmt.Printf(" %d ", cell.neighborCount) // Number
			}
		}
		fmt.Println() // Newline for the next row
	}
}