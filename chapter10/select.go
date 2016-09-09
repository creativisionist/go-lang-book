package main

import (
  "fmt"
  "time"
)

func main() {
  c1 := make(chan string)
  c2 := make(chan string)
  // c := make(chan int, 1)
  // Creates buffered channel with capacity of 1.
  // Normally channels are synchronous; both sides of the
  // channel will wait until the other side is ready. A
  // buffered channel is asynchronous; sending or receiving
  // a message will not wait unless the channel is already
  // full.

  go func() {
    for {
      c1 <- "from 1"
      time.Sleep(time.Second * 2)
    }
  }()
  go func() {
    for {
      c2 <- "from 2"
      time.Sleep(time.Second * 3)
    }
  }()
  go func() {
    for {
      select {
      case msg1 := <- c1:
        fmt.Println("Message 1", msg1)
      case msg2 := <- c2:
        fmt.Println("Message 2", msg2)
      case <- time.After(time.Second):
        fmt.Println("timeout")
      default:
        fmt.Println("nothing ready")
      }
    }
  }()
  var input string
  fmt.Scanln(&input)
}