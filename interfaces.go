package main

import ("fmt"; "math")

// A struct is a type which contains named fields. For example we could represent a Circle like this:
type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

type Shape interface {
    area() float64
    perimeter() float64
}

// Interfaces can also be used as fields:
type MultiShape struct {
    shapes []Shape
}

// We can even turn MultiShape itself into a Shape by giving it an area method:
func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
      area += s.area()
  }

  return area
}

func distance(x1, y1, x2, y2 float64) float64 {
  x := x2 - x1
  y := y2 - y1
  return math.Sqrt(x*x + y*y)
}

// Methods: ! Define methods inside struct

// In between the keyword func and the name of the function we've added a “receiver”.
// The receiver is like a parameter – it has a name and a type – but by creating
// the function in this way it allows us to call the function using the . operator:
func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
  return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func (r *Rectangle) perimeter() float64 {
  return 2 * ( (r.x2 - r.x1) + (r.y2 - r.y1) )
}

func main() {
  c := Circle{x: 0, y: 0, r: 5}
  fmt.Println(c.area())
  fmt.Println(c.perimeter())

  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r.area())
  fmt.Println(r.perimeter())

  fmt.Println(totalArea(&c, &r))
}
