package snake

import (
	"github.com/zacpez/scape-goat/api"
)

// FindFoodDirection removes bad ideas for the snake
func FindFoodDirection(snake *api.Snake, board *api.Board, direction *api.Direction) api.Direction {
	if len(board.Food) > 0 {
		for _, food := range board.Food {
			if *direction == api.DOWN && food.X > snake.Body[0].X {
				return api.DOWN
			}
			if *direction == api.UP && food.X < snake.Body[0].X {
				return api.UP
			}
			if *direction == api.RIGHT && food.Y > snake.Body[0].Y {
				return api.RIGHT
			}
			if *direction == api.LEFT && food.Y < snake.Body[0].Y {
				return api.LEFT
			}
		}
	}

	return api.NONE
}
