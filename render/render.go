package render

import (
	"bufio"
	"game_snake/game_board"
	"game_snake/snake"
	"os"
	"strings"
	"game_snake/point"
	"game_snake/food"
)


func Render(s *snake.Snake, b *board.Board, f food.Food) {
	writer := bufio.NewWriter(os.Stdout)
	str := ""
	for i := 0; i <= 10; i+=1 {
		str = str + "\n"
	}
	str = str + strings.Repeat("#", b.Width + 1)
	str = str + "\n"
	writer.WriteString(str)
	for i := 0; i < b.Height; i++ {
		str := []byte{'#'}
		for j := 0; j < b.Width; j++ {
			dot := point.Point{X : i, Y : j}
			if (j == b.Width - 1) {
				str = append(str, '#')
			} else {
				_, ok := s.Hashed_body[snake.Hash(dot)]
				if !ok {
					if dot.IsEqual(s.GetHeadCoord()) {
						str = append(str, 'O')
					} else {
						str = append(str, 'o')
					}
				} else if (dot ==  f.Coord) {
					str = append(str, '*')
				} else {
					str = append(str, ' ')
				}
			}
		}
		str = append(str, '\n')
		writer.Write(str)
	}
	str = strings.Repeat("#", b.Width + 1)
	str = str + "\n"
	writer.WriteString(str)
	writer.Flush()
}