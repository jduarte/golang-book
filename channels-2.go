package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            c1 <- "from 1"
            time.Sleep(time.Second * 1)
        }
    }()

    go func() {
        for {
            c2 <- "from 2"
            time.Sleep(time.Second * 5)
        }
    }()

    // go func() {
    //     for {
    //         // select picks the first channel that is ready and receives from it (or sends to it).
    //         // If more than one of the channels are ready then it randomly picks which one to receive from.
    //         // If none of the channels are ready, the statement blocks until one becomes available.
    //         select {
    //         case msg1 := <- c1:
    //             fmt.Println(msg1)
    //         case msg2 := <- c2:
    //             fmt.Println(msg2)
    //         }
    //     }
    // }()

    // The select statement is often used to implement a timeout:
    select {
    case msg1 := <- c1:
        fmt.Println("Message 1", msg1)
    case msg2 := <- c2:
        fmt.Println("Message 2", msg2)
    case <- time.After(time.Second):
        fmt.Println("timeout")
    }

    var input string
    fmt.Scanln(&input)
}
