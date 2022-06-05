package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver Square) End() Point {
	return Point{x: receiver.start.y, y: receiver.start.x}
}

func (receiver Square) Area() uint {
	return receiver.a * receiver.a
}

func (receiver Square) Perimeter() uint {
	return 4 * receiver.a
}
