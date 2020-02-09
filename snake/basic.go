package snake

import (
	"time"

	"github.com/zacpez/scape-goat/api"
)

// Difference of two Sets: A - B
func Difference(a []api.Direction, b []api.Direction) (diff []api.Direction) {
	m := make(map[api.Direction]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

// EdgeDirection finds a Direction if next to the edge, otherwise api.NONE
func EdgeDirection(snake *api.Snake, boardWidth int, boardHeight int) api.Direction {
	if snake.Body[0].X == 0 {
		return api.LEFT
	}
	if snake.Body[0].X == boardWidth-1 {
		return api.RIGHT
	}
	if snake.Body[0].Y == 0 {
		return api.UP
	}
	if snake.Body[0].Y == boardHeight-1 {
		return api.DOWN
	}
	return api.NONE
}

// NeckDirection finds the neck of a api.Snake
func NeckDirection(snake *api.Snake) api.Direction {
	var dx = snake.Body[0].X - snake.Body[1].X
	var dy = snake.Body[0].Y - snake.Body[1].Y
	if dx > 0 {
		return api.LEFT
	}
	if dx < 0 {
		return api.RIGHT
	}
	if dy < 0 {
		return api.DOWN
	}
	return api.UP
}

// SimpleRandomChoice gets an index given maxChoice indexes
func SimpleRandomChoice(maxChoice int) int {
	return (int(time.Now().UnixNano()) % maxChoice)
}
