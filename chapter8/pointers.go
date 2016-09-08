package main

import "fmt"

// func zero(x int) {
//   x = 0
// }
// func main() {
//   x := 5
//   zero(x)
//   fmt.Println(x)  // x is 5
// }

// Pointers reference a location in memory where a value
// is stored rather than the value itself. (They point to
// something else) By using a pointer (*int) the zero
// function is able to modify the original variable.
func zero(xPtr * int) {
  *xPtr = 0
}
func main() {
  x := 5
  zero(&x)
  fmt.Println(x) // x is 0
}