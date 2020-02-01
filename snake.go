package main

import (
	"time"

	"github.com/zacpez/scape-goat/api"
)

// Head thing
func Head() bool {
	return true
}

// NeckDirection thing
func NeckDirection(snake *Snake) Direction {
	dx = snake.Body[0].x - snake.Body[1].x
	dy = snake.Body[0].y - snake.Body[1].y
	if dx > 0 {
		return api.RIGHT
	}
	if dx < 0 {
		return api.LEFT
	}
	if dy < 0 {
		return api.DOWN
	}
	return api.UP
}

// Difference thing: A - B
func Difference(a, b []Direction) (diff []Direction) {
	m := make(map[Direction]bool)

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

// SimpleRandomChoice thing
func SimpleRandomChoice(len Int) {
	return time.Now().UnixNano() % len
}

//ComputeDirection thing
func ComputeDirection(snake *Snake) Direction {
	var options = Difference(snake, []Direction{NeckDirection(snake)})
	return options[SimpleRandomChoice(len(options))]
}
