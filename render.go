package main

import "strings"

func RenderArena(m *model) {
	m.arena = append(m.arena, strings.Split(m.verticalLine+strings.Repeat(m.horizontalLine, m.width)+m.verticalLine, ""))

	for i := 0; i < m.height; i++ {
		m.arena = append(m.arena, strings.Split(m.verticalLine+strings.Repeat(m.emptySymbol, m.width)+m.verticalLine, ""))
	}

	m.arena = append(m.arena, strings.Split(m.verticalLine+strings.Repeat(m.horizontalLine, m.width)+m.verticalLine, ""))

}

func RenderSnake(m *model) {
	for _, b := range m.snake.body {
		m.arena[b.x][b.y] = m.snakeSymbol
	}
}

func RenderFood(m *model) {
	m.arena[m.food.x][m.food.y] = m.foodSymbol
}
