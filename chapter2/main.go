package main

import "fmt"

// this is a comment
/* this is a much longer comment */
/* with multiple lines */

func main() {
  fmt.Println("Hello, my name is computer.")
  fmt.Println(len("Hello World"))
  fmt.Println("Hello World"[1])
  fmt.Println("Hello " + "World")

  // variable section
  var x string = "Hello World"
  fmt.Println(x)
  y := 5
  fmt.Println(y)
  f()
}

func f() {
  dogsName := "Max"
  fmt.Println("My dog's name is", dogsName)
}
