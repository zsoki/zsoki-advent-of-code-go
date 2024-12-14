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

var Directions = [8]Coord{DirN, DirNE, DirE, DirSE, DirS, DirSW, DirW, DirNW}

type Grid[T any] [][]T

func (grid *Grid[T]) get(coord Coord) T {
	return (*grid)[coord.Row][coord.Col]
}

func (grid *Grid[T]) set(coord Coord, value T) {
	(*grid)[coord.Row][coord.Col] = value
}
