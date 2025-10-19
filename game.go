package main

const (
	MINE  = -1
	EMPTY = 0
)

const (
	ONE   = iota + 1 // 1
	TWO              // 2
	THREE            // 3
	FOUR             // 4
	FIVE             // 5
	SIX              // 6
)

type Game struct {
	mine      int
	board     Board
	gameState int
}

type Board struct {
	mines [][]Mine
}

type Mine struct {
	visible       bool
	marked        bool
	neighborCount int
}

func main() {

}
