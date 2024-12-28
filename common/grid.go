package common

type Coord struct {
	Row, Col int
}

func (left Coord) Add(right Coord) Coord {
	return Coord{left.Row + right.Row, left.Col + right.Col}
}

func (left Coord) Sub(right Coord) Coord {
	return Coord{left.Row - right.Row, left.Col - right.Col}
}

func (left Coord) Gt(right Coord) bool {
	return left.Row > right.Row && left.Col > right.Col
}

func (left Coord) Lt(right Coord) bool {
	return left.Row < right.Row && left.Col < right.Col
}

func (left Coord) Eq(right Coord) bool {
	return left.Row == right.Row && left.Col == right.Col
}

func (left Coord) Times(mult int) Coord { return Coord{left.Row * mult, left.Col * mult} }

func (left Coord) CardinalNeighborsWhere(predicate func(neighbor Coord) bool) []Coord {
	trueNeighbors := make([]Coord, 0)
	for _, direction := range CardinalDirections {
		neighbor := left.Add(direction)
		if predicate(neighbor) {
			trueNeighbors = append(trueNeighbors, neighbor)
		}
	}
	return trueNeighbors
}

func SameLine(a Coord, b Coord) bool {
	return a.Row == b.Row || a.Col == b.Col
}

var (
	DirN  = Coord{-1, 0}
	DirNE = Coord{-1, 1}
	DirE  = Coord{0, 1}
	DirSE = Coord{1, 1}
	DirS  = Coord{1, 0}
	DirSW = Coord{1, -1}
	DirW  = Coord{0, -1}
	DirNW = Coord{-1, -1}
)

var AllDirections = [8]Coord{DirN, DirNE, DirE, DirSE, DirS, DirSW, DirW, DirNW}
var CardinalDirections = [4]Coord{DirN, DirE, DirS, DirW}

type Grid[T any] [][]T

func (grid *Grid[T]) get(coord Coord) T {
	return (*grid)[coord.Row][coord.Col]
}

func (grid *Grid[T]) set(coord Coord, value T) {
	(*grid)[coord.Row][coord.Col] = value
}
