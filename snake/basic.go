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
func EdgeDirection(snake *api.Snake, boardWidth int, boardHeight int) []api.Direction {
	var edges = []api.Direction{}
	if snake.Body[0].X == 0 {
		edges = append(edges, api.LEFT)
	}
	if snake.Body[0].X == boardWidth-1 {
		edges = append(edges, api.RIGHT)
	}
	if snake.Body[0].Y == 0 {
		edges = append(edges, api.UP)
	}
	if snake.Body[0].Y == boardHeight-1 {
		edges = append(edges, api.DOWN)
	}
	return edges
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

// SimpleAvoidance is just a one level check for obsticales
func SimpleAvoidance(snake *api.Snake, board *api.Board, direction *api.Direction) api.Direction {
	var dx = snake.Body[0].X
	var dy = snake.Body[0].Y
	if *direction == api.DOWN {
		dy = dy + 1
	}
	if *direction == api.UP {
		dy = dy - 1
	}
	if *direction == api.LEFT {
		dx = dx - 1
	}
	if *direction == api.RIGHT {
		dx = dx + 1
	}

	for _, other := range board.Snakes {
		for _, part := range other.Body {
			if part.X == dx && part.Y == dy {
				return *direction
			}
		}
	}
	return api.NONE
}

// SimpleRandomChoice gets an index given maxChoice indexes
func SimpleRandomChoice(maxChoice int) int {
	return (int(time.Now().UnixNano()) % maxChoice)
}
