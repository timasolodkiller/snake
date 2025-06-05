package snake

import (
	"game_snake/game_board"
	"game_snake/point"
)

type Snake struct {
	body        *queue
	size        int
	Hashed_body map[uint64]int
	Free_points []point.Point
}

func Hash(p point.Point) uint64 {
	tmp := p.X*23 + p.Y*7
	return uint64(tmp % 5000)
}
func (s *Snake) GetHeadCoord() point.Point {
	return s.body.head.next.coord
}
func (s *Snake) GetHeadDirection() Direction {
	return s.body.head.next.direction
}

func Initialize(p point.Point, direction Direction, b *board.Board) *Snake {
	k := 0
	m := make(map[uint64]int)
	f := make([]point.Point, (b.Width) * (b.Height))
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			if i != p.X || j != p.Y {
				p := point.Point{X: i, Y: j}
				m[Hash(p)] = k
				f[k] = p
				k += 1
			}
		}
	}
	q := New()
	q.PushFront(p, direction)
	s := &Snake{
		body:        q,
		size:        1,
		Hashed_body: m,
		Free_points: f,
	}
	return s
}

func (s *Snake) GetCoordsByMoveHead(dir Direction) point.Point {
	var p point.Point
	switch dir {
	case Up:
		p = point.Point{X: s.body.head.next.coord.X - 1, Y: s.body.head.next.coord.Y}
	case Down:
		p = point.Point{X: s.body.head.next.coord.X + 1, Y: s.body.head.next.coord.Y}
	case Left:
		p = point.Point{X: s.body.head.next.coord.X, Y: s.body.head.next.coord.Y - 1}
	case Right:
		p = point.Point{X: s.body.head.next.coord.X, Y: s.body.head.next.coord.Y + 1}
	}
	return p
}

func (s *Snake) GetCoordsByMoveTail(dir Direction) point.Point {
	var p point.Point
	switch dir {
	case Up:
		p = point.Point{X: s.body.tail.coord.X - 1, Y: s.body.tail.coord.Y}
	case Down:
		p = point.Point{X: s.body.tail.coord.X + 1, Y: s.body.tail.coord.Y}
	case Left:
		p = point.Point{X: s.body.tail.coord.X, Y: s.body.tail.coord.Y - 1}
	case Right:
		p = point.Point{X: s.body.tail.coord.X, Y: s.body.tail.coord.Y + 1}
	}
	return p
}
func (s *Snake) Move(direction Direction) bool {
	next_point := s.GetCoordsByMoveHead(direction)
	prev_point, _ := s.body.PopBack()
	prev_point_Hash := Hash(prev_point)
	next_point_Hash := Hash(next_point)
	s.Hashed_body[prev_point_Hash] = s.Hashed_body[next_point_Hash]
	s.Free_points[s.Hashed_body[next_point_Hash]] = prev_point
	if !s.IsIntersect(next_point) {
		return false
	}
	delete(s.Hashed_body, next_point_Hash)
	if s.size > 1 {
		s.body.head.next.direction = direction
	}
	s.body.PushFront(next_point, direction)
	return true
}

func (s *Snake) Grow() {
	direction := s.body.tail.direction
	switch direction {
	case Up:
		direction = Down
	case Down:
		direction = Up
	case Left:
		direction = Right
	case Right:
		direction = Left
	}
	p := s.GetCoordsByMoveTail(direction)
	s.body.PushBack(p, s.body.tail.direction)
	p_Hash := Hash(p)
	_, ok := s.Hashed_body[p_Hash]
	if ok {
		s.Free_points[s.Hashed_body[p_Hash]] = s.Free_points[len(s.Free_points)-1]
		s.Hashed_body[Hash(s.Free_points[len(s.Free_points)-1])] = s.Hashed_body[p_Hash]
		s.Free_points = s.Free_points[:len(s.Free_points)-1]
		delete(s.Hashed_body, p_Hash)
	}
	s.size += 1
}

func (s *Snake) PrintFromHead() {
	s.body.PrintFromHead()
}

func (s *Snake) PrintFromTail() {
	s.body.PrintFromTail()
}

func (s *Snake) IsIntersect(p point.Point) bool {
	_, ok := s.Hashed_body[Hash(p)]
	return ok
}
