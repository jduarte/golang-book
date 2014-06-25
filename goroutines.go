package main

import (
    "fmt"
    "time"
    "math/rand"
)

// func f1(n int) {
//     for i := 0; i < 10; i++ {
//         fmt.Println(n, ":", i)
//     }
// }

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    go f(0)
    var input string
    // With a goroutine we return immediately to the next line and don't wait for the function to complete.
    // This is why the call to the Scanln function has been included; without it the program would exit before being given the opportunity to print all the numbers.
    // fmt.Scanln(&input)

    // Goroutines are lightweight and we can easily create thousands of them. We can modify our program to run 10 goroutines by doing this:
    // for i := 0; i < 10; i++ {
    //     go f1(i)
    // }
    // fmt.Scanln(&input)

    // f prints out the numbers from 0 to 10, waiting between 0 and 250 ms after each one. The goroutines should now run simultaneously.
    for i := 0; i < 10; i++ {
        go f(i)
    }
    fmt.Scanln(&input)
}
