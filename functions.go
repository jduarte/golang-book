package main

import "fmt"

func average(xs []float64) float64 {
  total := 0.0

  if len(xs) == 0 {
    return total
  }

  for _, v := range xs {
    total = total + v
  }

  return total / float64(len(xs))
}

// We can return multiple values
func mult_values_return() (int, int) {
  return 5, 9
}

// Varidadic functions, multiple arguments
func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }

  return total
}

// Closures
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

// Recursion
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }

    return x * factorial(x - 1)
}

func main() {
  xs := []float64 { 98, 93, 77, 82, 83 }
  fmt.Println(average(xs))

  x, y := mult_values_return()
  fmt.Println(x, y)

  // we can add as many arguments as we want
  fmt.Println(add(1,2,3))
  // we can also pass slices to it
  ss := []int{1,2,3}
  fmt.Println(add(ss...))


    // Closures.
        // You can create functions within functions
    add_closure := func(x, y int) int {
        return x + y
    }
    fmt.Println(add_closure(1,1))

        // One way to use closure is by writing a function which returns another function which – when called – can generate a sequence of numbers.
        // For example here's how we might generate all the even numbers:

    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4

    // Recursion
    fmt.Println(factorial(5))
}
