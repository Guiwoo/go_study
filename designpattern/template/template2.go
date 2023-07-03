package template

import "fmt"

/**
Algorithms can be decomposed into common parts + specifics
Strategy pattern does this through composition
	- High-level algorithm uses an interface
	- Concrete implementations implement the interface
	- We keep a pointer to the interface; provide concrete implementations

Template Method performs a similar operation, but
	- It's typically just a function, not a struct with a reference to the implementation
	- Can still use interface ; or
	- Can be functional (take several functions as parameters

A skeleton algorithm defined in a function.
Function can either use an interface or can take several functions as arguments.
*/

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins. \n", g.WinningPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func (c *chess) Start() {
	fmt.Println("Starting a new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

var _ Game = (*chess)(nil)

func PlayGame2(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}

func Start2() {
	chess := NewGameOfChess()
	PlayGame(chess)
}
