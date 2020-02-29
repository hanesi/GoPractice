package main

import "fmt"
import "math"

type circle struct {
	rad float64
}

type square struct {
	L float64
}

func (c circle) area() float64 {
	return c.rad * c.rad * math.Pi
}

func (s square) area() float64 {
	return s.L * s.L
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())

}

func main() {
	circ := circle{
		rad: 12.345,
	}

	sq := square{
		L: 45.578,
	}

	info(circ)
	info(sq)
}
