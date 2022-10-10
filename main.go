package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	Left = 1 + iota
	Right
	Up
	Down
)

const (
	HELP      = "\n\nUse arrow keys or w a s d keys to nagivate the snake.\n"
	GAME_OVER = "\n\nGAME OVER.\n"
	QUIT      = "\n\nPress q to quit.\n"
)

const (
	INTERVAL = 80
)

type model struct {
	stopwatch      stopwatch.Model
	horizontalLine string
	verticalLine   string
	emptySymbol    string
	snakeSymbol    string
	foodSymbol     string
	width          int
	height         int
	arena          [][]string
	snake          snake
	lostGame       bool
	score          int
	food           food
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		stopwatch:      stopwatch.NewWithInterval(time.Duration(INTERVAL) * time.Millisecond),
		horizontalLine: "-",
		verticalLine:   "|",
		emptySymbol:    " ",
		snakeSymbol:    "#",
		foodSymbol:     "$",
		width:          40,
		height:         20,
		arena:          [][]string{},
		lostGame:       false,
		score:          0,
		food: food{
			x: 10, y: 20,
		},
		snake: snake{
			body: []foodLocation{{x: 1, y: 1},
				{x: 1, y: 2},
				{x: 1, y: 3},
				{x: 1, y: 4}},
			length:    4,
			direction: Right,
		},
	}
}

func (m model) Init() tea.Cmd {
	var x, y int

	x = rand.Intn(m.height)
	y = rand.Intn(m.width)

	m.food.x = x
	m.food.y = y
	return m.stopwatch.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "w":
			if m.snake.hitWall(m) {
				m.lostGame = true
				return m, tea.Quit
			}

			m.snake.direction = Up

			return m, nil

		case "down", "s":
			if m.snake.hitWall(m) {
				m.lostGame = true
				return m, tea.Quit
			}

			m.snake.direction = Down

			return m, nil

		case "left", "a":
			if m.snake.hitWall(m) {
				m.lostGame = true
				return m, tea.Quit
			}

			m.snake.direction = Left

			return m, nil
		case "right", "d":
			if m.snake.hitWall(m) {
				m.lostGame = true
				return m, tea.Quit
			}
			m.snake.direction = Right

			return m, nil

		}

	}
	var cmd tea.Cmd
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	h := m.snake.getHead()
	c := foodLocation{x: h.x, y: h.y}

	switch m.snake.direction {
	case Right:
		c.y++
	case Left:
		c.y--
	case Up:
		c.x--
	case Down:
		c.x++
	}

	if c.x == m.food.x && c.y == m.food.y {
		m.snake.length++
		x := rand.Intn(m.height - 1)
		y := rand.Intn(m.width - 1)

		for {
			if !m.snake.hitSelf(foodLocation{x, y}) {
				break
			}
		}

		m.food.x = x
		m.food.y = y

	}

	if m.snake.hitWall(m) || m.snake.hitSelf(c) {
		m.lostGame = true
		return m, tea.Quit
	}

	if len(m.snake.body) < m.snake.length {
		m.snake.body = append(m.snake.body, c)
		m.score = m.score + 10

	} else {
		m.snake.body = append(m.snake.body[1:], c)
	}

	return m, cmd

}

func (m model) View() string {
	// The header
	s := "Go Snake!!!\n\n"

	stringArena := ""

	m.arena = append(m.arena, strings.Split(m.verticalLine+strings.Repeat(m.horizontalLine, m.width)+m.verticalLine, ""))

	for i := 0; i < m.height; i++ {
		m.arena = append(m.arena, strings.Split(m.verticalLine+strings.Repeat(m.emptySymbol, m.width)+m.verticalLine, ""))
	}

	m.arena = append(m.arena, strings.Split(m.verticalLine+strings.Repeat(m.horizontalLine, m.width)+m.verticalLine, ""))

	for _, b := range m.snake.body {
		m.arena[b.x][b.y] = m.snakeSymbol
	}

	m.arena[m.food.x][m.food.y] = m.foodSymbol

	for _, row := range m.arena {
		stringArena += strings.Join(row, "") + "\n"
	}

	s += stringArena

	s += fmt.Sprintf("\n\nScore: %d", m.score)

	if m.lostGame {
		s += "\n\nGame Over.\n"
	}

	s += HELP

	// The footer
	s += "\nPress q to quit.\n"

	return s
}
