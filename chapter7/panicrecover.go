package main

import "fmt"

// Call to recover will never happen b/c call to panic
// immediately stops execution of the function.
// func main() {
//   panic("PANIC")
//   str := recover()
//   fmt.Println(str)
// }

// Instead we pair it with defer:
func main() {
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}