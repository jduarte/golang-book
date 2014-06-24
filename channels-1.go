package main

import (
    "fmt"
    "time"
)

// Channels provide a way for two goroutines to communicate with one another and synchronize their execution. Here is an example program using channels:
func pinger(c chan string) {
    for i := 0; ; i++ {
        // The <- (left arrow) operator is used to send and receive messages on the channel. c <- "ping" means send "ping"
        c <- "ping"
    }
}

func ponger(c chan string) {
  for i := 0; ; i++ {
        c <- "pong"
    }
}

func printer(c chan string) {
    for {
        fmt.Println(<- c) // means receive a message from channel 'c' and print it
        time.Sleep(time.Second * 1)
    }
}

// Channel Direction
  // Now c can only be sent to. Attempting to receive from c will result in a compiler error. Similarly we can change printer to this:
  // func printer(c chan<- string)

  // func printer(c <-chan string)

  // A channel that doesn't have these restrictions is known as bi-directional. A bi-directional channel can be passed to a function that takes send-only or receive-only channels, but the reverse is not true.
func main() {
    // A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel
    var c chan string = make(chan string)
    
    // Using a channel like this synchronizes the two goroutines. When pinger attempts to send a message on the channel
    // it will wait until printer is ready to receive the message. (this is known as blocking)
    go pinger(c)
    go ponger(c)
    go printer(c)
    
    var input string
    fmt.Scanln(&input)
}
