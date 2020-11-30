package graph

type Point2D struct {
	X int
	Y int
}

type Point3D struct {
	*Point2D
	Z int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var Directions []Direction = []Direction{North, East, South, West}

func CardinalDirections() []Direction {
	return Directions
}

func Move(p Point2D, d Direction) Point2D {
	if d == North {
		return Point2D{X: p.X, Y: p.Y - 1}
	} else if d == East {
		return Point2D{X: p.X + 1, Y: p.Y}
	} else if d == South {
		return Point2D{X: p.X, Y: p.Y + 1}
	} else if d == West {
		return Point2D{X: p.X - 1, Y: p.Y}
	}

	// unknown direction
	return Point2D{}
}

func Adj(p Point2D) []Point2D {
	adjPoint := make([]Point2D, 0, len(Directions))
	for _, dir := range Directions {
		a := Move(p, dir)
		adjPoint = append(adjPoint, a)
	}
	return adjPoint
}
