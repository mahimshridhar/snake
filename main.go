package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

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

func main() {
	rand.Seed(time.Now().UnixNano())
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
