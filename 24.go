package main

import "math"

func main() {
	println(calculateDistance(NewPoint(23, -11), NewPoint(-1, 42)))
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func calculateDistance(first, second Point) float64 {
	lenX := second.x - first.x
	lenY := second.y - first.y

	distance := math.Sqrt(math.Pow(lenX, 2) + math.Pow(lenY, 2))

	return distance
}
