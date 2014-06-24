package main

import ("fmt"; "math")

// A struct is a type which contains named fields. For example we could represent a Circle like this:
type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  x := x2 - x1
  y := y2 - y1
  return math.Sqrt(x*x + y*y)
}

type Person struct {
    Name string
}

func (p* Person) Talk() {
  fmt.Println("My name is:", p.Name)
}

// A struct's fields usually represent the has-a relationship. For example a Circle has a radius. Suppose we had a person struct:
type Android struct {
    Person Person
    Model string
} // This would work, but we would rather say an Android is a Person, rather than an Android has a Person.

// Go supports relationships like this by using an embedded type.
// Also known as anonymous fields, embedded types look like this:
type AndroidEmbeddedType struct {
  Person // We use the type (Person) and don't give it a name. When defined this way the Person struct can be accessed using the type name:
  Model string
}

// Methods: ! Define methods inside struct

// In between the keyword func and the name of the function we've added a “receiver”.
// The receiver is like a parameter – it has a name and a type – but by creating
// the function in this way it allows us to call the function using the . operator:
func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func main() {
  c := Circle{x: 0, y: 0, r: 5}
  fmt.Println(c.area())

  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r.area())

  // Embedded Types
    // Since we then defined Person as an anonymous field in AndroidEmbeddedType, the latter class automatically can call on all the visible behaviors/methods of the anonymous field type.
    // So here, we have not subclassed a parent class, but composed it. But the effect is the very same

    // But we can also call any Person methods directly on the Android:
    a1 := new(Android)
    a1.Person.Talk()

    a2 := new(AndroidEmbeddedType)
    // But we can also call any Person methods directly on the Android:
    a2.Talk()
    // The is-a relationship works this way intuitively: People can talk, an android is a person, therefore an android can talk.

}
