package point

import "strconv"

type Point struct {
	X, Y int 
}

func (p Point) String() string{
	return strconv.Itoa(p.X) + " " + strconv.Itoa(p.Y)
}

func (p Point) IsEqual(dup Point) bool {
	return (p.X == dup.X && p.Y == dup.Y)
}