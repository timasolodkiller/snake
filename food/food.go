package food

import (
	"game_snake/point"
	"game_snake/snake"
	"math/rand"
	"os"
	"strconv"
)

type Food struct {
	Coord point.Point
	Rand *rand.Rand
}

func Initialize(s *snake.Snake, r *rand.Rand, f *os.File) Food {
	index := r.Intn(len(s.Free_points) - 1) + 1
	text := strconv.Itoa(s.Free_points[index].X) + " " + strconv.Itoa(s.Free_points[index].Y) + " " +  strconv.Itoa(index) + "\n"
	f.WriteString(text)
	return Food{
		Coord : s.Free_points[index],
		Rand : r,
	}
}
