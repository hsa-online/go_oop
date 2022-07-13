package main 

import (
  "fmt"
  "math"
)

type Shape struct {
  instance interface{}

  IsLargerThan func(this interface{}, area float64) bool
  GetArea func(this interface{}) float64
}

func NewShape(instance interface{}) *Shape {
  this := &Shape{}
  this.instance = instance
  this.IsLargerThan = ShapeIsLargerThan

  return this
}

func ShapeIsLargerThan(this interface{}, area float64) bool {
  if shape, ok := this.(*Shape); ok {
    return shape.GetArea(shape.instance) > area
  } else {
    panic(fmt.Errorf("wrong type passed %T", this))
  }
}

type Circle struct {
  parent *Shape

  radius float64

  IsLargerThan func(this interface{}, area float64) bool
  GetArea func(this interface{}) float64
}

func NewCircle(radius float64) *Circle {
  this := &Circle{}
  this.parent = NewShape(this)
  this.parent.GetArea = CircleGetArea
  this.radius = radius
  this.IsLargerThan = CircleIsLargerThan
  this.GetArea = CircleGetArea

  return this
}

func (circle Circle) String() string {
  return fmt.Sprintf("Circle (radius=%f)", circle.radius)
}

func CircleIsLargerThan(this interface{}, area float64) bool {
  if circle, ok := this.(*Circle); ok {
    return circle.parent.IsLargerThan(circle.parent, area)
  } else {
    panic(fmt.Errorf("wrong type passed %T", this))
  }
}

func CircleGetArea(this interface{}) float64 {
  if circle, ok := this.(*Circle); ok {
    return math.Pi * circle.radius * circle.radius
  } else {
    panic(fmt.Errorf("wrong type passed %T", this))
  } 
}

type Square struct {
  parent *Shape

  side float64

  IsLargerThan func(this interface{}, area float64) bool
  GetArea func(this interface{}) float64
}

func NewSquare(side float64) *Square {
  this := &Square{}
  this.parent = NewShape(this)
  this.parent.GetArea = SquareGetArea
  this.side = side
  this.IsLargerThan = SquareIsLargerThan 
  this.GetArea = SquareGetArea

  return this
}

func (square Square) String() string {
  return fmt.Sprintf("Square (side=%f)", square.side)
}

func SquareIsLargerThan(this interface{}, area float64) bool {
  if square, ok := this.(*Square); ok {
    return square.parent.IsLargerThan(square.parent, area)
  } else {
    panic(fmt.Errorf("wrong type passed %T", this))
  }
}

func SquareGetArea(this interface{}) float64 {
  if square, ok := this.(*Square); ok {
    return square.side * square.side
  } else {
    panic(fmt.Errorf("wrong type passed %T", this))
  } 
}

func main() {
  area := 16.0

  square := NewSquare(5)
  fmt.Printf("%v is larger than %f: %t\n", square, area, square.IsLargerThan(square, area))

  circle := NewCircle(2)
  fmt.Printf("%v is larger than %f: %t\n", circle, area, circle.IsLargerThan(circle, area))
} 
