package main

import ("fmt"; "math")


// A struct is a type which contains named fields. For example we could represent a Circle like this:
type Circle struct {
    x float64
    y float64
    r float64
}

func circleArea(c Circle) float64 {
  return math.Pi * c.r * c.r
}

// One thing to remember is that arguments are always copied in Go.
// If we attempted to modify one of the fields inside of the circleArea function, it would not modify
// the original variable. Because of this we would typically write the function like this:
func circleAreaPtr(c* Circle) float64 {
  return math.Pi * c.r * c.r
}

func main() {
  // Like with other data types, this will create a local Circle variable that is by default set to zero. For a struct zero means each of the fields is set to their corresponding zero value (0 for ints, 0.0 for floats, "" for strings, nil for pointers, …)
  //var c Circle

  // This allocates memory for all the fields, sets each of them to their zero value and returns a pointer. (*Circle)
  //c1 := new(Circle)

  // More often we want to give each of the fields a value.
  c2 := Circle{x: 0, y: 0, r: 5}

  // Or we can leave off the field names if we know the order they were defined:
  c3 := Circle{0, 0, 5}

  fmt.Println(c2.x, c2.y, c2.r)


  fmt.Println(circleArea(c2))

  fmt.Println(circleAreaPtr(&c3))
}
