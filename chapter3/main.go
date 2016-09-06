package main

import "fmt"

func main() {
  fmt.Println("1 + 1 =", 1 + 1)
  fmt.Println("____")
  fmt.Println("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}