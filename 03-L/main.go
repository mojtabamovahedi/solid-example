package main

import (
	"fmt"
)

func main() {
	rect := Rectangle{}
	rect.SetHeight(10)
	rect.SetWidth(5)
	fmt.Println(rect.area())

	square := Square{}
	square.SetWidth(12)
	fmt.Println(square.area())
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

func (r *Rectangle) SetWidth(w float64) {
	r.width = w
}

func (r *Rectangle) SetHeight(h float64) {
	r.height = h
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

type Square struct {
	Rectangle
}

func (s *Square) SetHeight(h float64) {
	s.height = h
	s.width = h
}

func (s *Square) SetWidth(w float64) {
	s.width = w
	s.height = w
}

func (s Square) area() float64 {
	return s.height * s.width
}

// we know square is a rectangle

//type Square struct {
//	side float64
//}
//
//func (s *Square) SetSide(side float64) {
//	s.side = side
//}
//
//func (s Square) area() float64 {
//	return s.side * s.side
//}
//
//type Shape interface {
//	area() float64
//}
