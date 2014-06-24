package main

import "fmt"

func zero(x int) {
    x = 0
}

func onePtr(xPtr *int) {
  *xPtr = 1
}

// Pointers reference a location in memory where a value is stored rather than the value itself.
func zeroPtr(xPtr *int) {
    *xPtr = 0
}

func square(xPtr *int) {
  *xPtr = *xPtr * *xPtr
}

func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

func main() {
    x := 5
    zero(x)
    fmt.Println(x) // x is still 5

    zeroPtr(&x)
    fmt.Println(x) // x is now 0

    // * is also used to “dereference” pointer variables. Dereferencing a pointer gives us access to the
    // value the pointer points to.

    // When we write *xPtr = 0 we are saying “store the int 0 in the memory location xPtr refers to”.
    // Dereferencing (*) a pointer gives us access to the value the pointer points to.

    // If we try xPtr = 0 instead we will get a compiler error because xPtr is not an int it's a *int, which can only be given another *int.

    // Finally we use the & operator to find the address of a variable. &x returns a *int (pointer to an int) because x is an int. This is what allows us to modify the original variable. &x in main and xPtr in zero refer to the same memory location.



    // Another way to get a pointer is to use the built-in new function:
    intPtr := new(int)
    onePtr(intPtr)
    fmt.Println(*intPtr) // 1

    y := 2
    square(&y)
    fmt.Println(y)

    x1 := 1
    x2 := 2

    swap(&x1, &x2)
    fmt.Println("x1:", x1, "x2:", x2)
}

