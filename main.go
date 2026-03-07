package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/term"
)

type coord struct {
	x int
	y int
}

type Game struct {
	snake   []coord
	food    coord
	score   int
	width   int
	height  int
	dir     int
	running bool
	mu      sync.Mutex
}

func (g *Game) spawnFood() {
	for {
		food := coord{
			x: rand.Intn(g.height-2) + 1,
			y: rand.Intn(g.width-2) + 1,
		}

		onSnake := false

		for _, part := range g.snake {
			if part == food {
				onSnake = true
				break
			}
		}

		if onSnake != true {
			g.food = food
			break
		}
	}
}

func (g *Game) readInput() {
	buf := make([]byte, 3)
	for g.running {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			continue
		}

		g.mu.Lock()
		switch {
		case n == 3 && buf[0] == 27 && buf[1] == '[':
			switch buf[2] {
			case 'A':
				if g.dir != 3 {
					g.dir = 1
				}
			case 'B':
				if g.dir != 1 {
					g.dir = 3
				}
			case 'C':
				if g.dir != 4 {
					g.dir = 2
				}
			case 'D':
				if g.dir != 2 {
					g.dir = 4
				}
			}

		case n == 1:
			switch buf[0] {
			case 'w', 'W':
				if g.dir != 3 {
					g.dir = 1
				}
			case 's', 'S':
				if g.dir != 1 {
					g.dir = 3
				}
			case 'a', 'A':
				if g.dir != 2 {
					g.dir = 4
				}
			case 'd', 'D':
				if g.dir != 4 {
					g.dir = 2
				}
			case 'q', 'Q':
				g.running = false
			}
		}
		g.mu.Unlock()
	}
}

func (g *Game) move() {
	g.mu.Lock()
	dir := g.dir
	g.mu.Unlock()

	head := g.snake[len(g.snake)-1]
	newHead := head

	switch dir {
	case 1:
		newHead.x--
	case 2:
		newHead.y++
	case 3:
		newHead.x++
	case 4:
		newHead.y--
	}

	if newHead.x <= 0 || newHead.x >= g.height-1 ||
		newHead.y <= 0 || newHead.y >= g.width-1 {
		g.running = false
		return
	}

	for i := 0; i < len(g.snake)-1; i++ {
		if g.snake[i] == newHead {
			g.running = false
			return
		}
	}

	g.snake = append(g.snake, newHead)

	if newHead == g.food {
		g.score += 10
		g.spawnFood()

	} else {
		g.snake = g.snake[1:]
	}

}

func (g *Game) draw() {
	grid := make([][]rune, g.height)

	for i := range grid {
		grid[i] = make([]rune, g.width)

		for j := range grid[i] {
			grid[i][j] = ' '
		}

	}

	for j := 0; j < g.width; j++ {
		grid[0][j] = '#'
		grid[g.height-1][j] = '#'
	}

	for i := 0; i < g.height; i++ {
		grid[i][0] = '#'
		grid[i][g.width-1] = '#'
	}

	for i, part := range g.snake {
		if i == len(g.snake)-1 {
			grid[part.x][part.y] = '@'
		} else {
			grid[part.x][part.y] = 'o'
		}
	}

	grid[g.food.x][g.food.y] = '*'

	var sb strings.Builder

	sb.WriteString("\033[H")

	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			sb.WriteRune(grid[i][j])
		}
		sb.WriteString("\r\n")

	}
	sb.WriteString(fmt.Sprintf("\r\nScore: %d\r\n", g.score))
	fmt.Print(sb.String())
}

func setupTerminal() func() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	fmt.Print("\033[?25l")

	return func() {
		fmt.Print("\033[?25h")
		term.Restore(int(os.Stdin.Fd()), oldState)
	}
}

func newGame() *Game {
	g := &Game{
		width:   50,
		height:  20,
		dir:     2,
		running: true,
	}

	midY := g.height / 2

	g.snake = []coord{
		{x: midY, y: 4},
		{x: midY, y: 5},
		{x: midY, y: 6},
	}
	g.spawnFood()

	return g

}

func main() {
	defer setupTerminal()()

	fmt.Print("\033[2J")

	game := newGame()

	go game.readInput()

	ticker := time.NewTicker(150 * time.Millisecond)

	defer ticker.Stop()

	for game.running {
		<-ticker.C

		game.move()
		game.draw()
	}

	fmt.Print("\033[2J\033[H")
	fmt.Printf("\r\nGAME OVER!  Final score: %d\r\n", game.score)
}
