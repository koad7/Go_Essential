// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"log"
)

type Square struct {
	Center Point
	Length int
}
type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}

	s := &Square{
		Center: Point{x,y}
		Length: length,
	}
	return s, nil
}


func (s *Square) Area() int {
	return s.Length * s.Length
}

func main(){
	s, err := NewSquare(1,1,10)
	if err != nil{
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2,3)
	fmt.Printf("%+v\n",s)
	fmt.Println(s.Area())
}