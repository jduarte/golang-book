package main

import "fmt"

func main() {
    // 1st Way
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83

    var total1 float64 = 0
    for i := 0; i < 5; i++ {
        total1 += x[i]
    }

    fmt.Println(total1 / 5)

    // 2nd way
    var total2 float64 = 0
    for i := 0; i < len(x); i++ {
        total2 += x[i]
    }

      // The issue here is that len(x) and total have different types. total is a float64 while len(x) is an int.
      // So we need to convert len(x) into a float64:
    fmt.Println(total2 / float64(len(x)))

    // 3rd way


    var total3 float64 = 0

      // In this for loop _ represents the current position in the array and value is the same as x[_].
      // We use the keyword range followed by the name of the variable we want to loop over.
      // A single _ (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator variable)
    for _, value := range x {
        total3 += value
    }
    fmt.Println(total3 / float64(len(x)))

    // Ways of declarating arrays
    var _ = [5]float64{ 98, 93, 77, 82, 83 }
    var _ = [4]float64{
        98,
        93,
        77,
        82,
        // 83,
    }

}
