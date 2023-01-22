package main

import (
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	INTERVAL = 80
)

type TickMsg time.Time

type Model struct {
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

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Second/10, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) changeSnakeDirection(direction int) (tea.Model, tea.Cmd) {
	if m.snake.hitWall(m) {
		m.lostGame = true
		return m, tea.Quit
	}

	opposites := map[int]int{
		Up:    Down,
		Down:  Up,
		Left:  Right,
		Right: Left,
	}

	if opposites[direction] != m.snake.direction {
		m.snake.direction = direction

	}

	return m, nil
}

func (m Model) moveSnake() (tea.Model, tea.Cmd) {
	h := m.snake.getHead()
	c := coord{x: h.x, y: h.y}

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
		x := rand.Intn(m.height-1) + 1
		y := rand.Intn(m.width-1) + 1

		for {
			if !m.snake.hitSelf(coord{x, y}) {
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
	return m, m.tick()
}

func initialModel() Model {
	return Model{
		horizontalLine: "#",
		verticalLine:   "#",
		emptySymbol:    " ",
		snakeSymbol:    "o",
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
			body: []coord{
				{x: 1, y: 1},
				{x: 1, y: 2},
				{x: 1, y: 3},
				{x: 1, y: 4}},
			length:    4,
			direction: Right,
		},
	}
}

func (m Model) Init() tea.Cmd {
	var x, y int

	x = rand.Intn(m.height - 1)
	y = rand.Intn(m.width - 1)

	m.food.x = x
	m.food.y = y
	return m.tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "w":
			return m.changeSnakeDirection(Up)

		case "down", "s":
			return m.changeSnakeDirection(Down)

		case "left", "a":
			return m.changeSnakeDirection(Left)

		case "right", "d":
			return m.changeSnakeDirection(Right)

		}
	case TickMsg:
		return m.moveSnake()

	}

	return m, nil

}

func (m Model) View() string {

	var sb strings.Builder

	sb.WriteString(RenderTitle())

	sb.WriteByte('\n')

	var stringArena strings.Builder

	RenderArena(&m)

	RenderSnake(&m)

	RenderFood(&m)

	for _, row := range m.arena {
		stringArena.WriteString(strings.Join(row, "") + "\n")
	}

	sb.WriteString(stringArena.String())
	sb.WriteByte('\n')

	sb.WriteString(RenderScore(m.score))
	sb.WriteByte('\n')

	if m.lostGame {
		sb.WriteString(RenderGameOver())

	}

	// sb.WriteString(fmt.Sprintf("\nx =%d\ny =%d\nsnake x=%d\nsnake y=%d\n", m.food.x, m.food.y, m.snake.getHead().x, m.snake.getHead().y))

	sb.WriteString(RenderHelp(HELP))
	sb.WriteByte('\n')

	// The footer
	sb.WriteString(RenderHelp("Press q or ctrl+c to quit."))
	// sb.WriteByte('\n')
	sb.WriteByte('\n')
	sb.WriteByte('\n')

	return sb.String()
}
