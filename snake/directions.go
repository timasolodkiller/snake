package snake

type Direction int 

const (
	Up Direction = iota
	Left
	Down
	Right  
)

func (m Direction) String() string{
	var res string
	switch m {
	case Up: 
		res = "Up"
	case Left: 
		res = "Left"
	case Right: 
		res = "Right"
	case Down: 
		res = "Down"
	}
	return res
}