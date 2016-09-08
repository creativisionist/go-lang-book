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
// func zero(xPtr * int) {
//   *xPtr = 0
// }
// func main() {
//   x := 5
//   zero(&x)
//   fmt.Println(x) // x is 0
// }

// Pointer is represented using the * (asterisk)
// character followed by the type of the stored value. In
// the zero function xPtr is a pointer to an int.
// * is also used to “dereference” pointer variables. Dereferencing
// a pointer gives us access to the value the
// pointer points to. When we write *xPtr = 0 we are saying
// “store the int 0 in the memory location xPtr refers
// to”. If we try xPtr = 0 instead we will get a compiler
// error because xPtr is not an int it's a *int, which can
// only be given another *int.
// Finally we use the & operator to find the address of a
// Pointers 94
// variable. &x returns a *int (pointer to an int) because x
// is an int. This is what allows us to modify the original
// variable. &x in main and xPtr in zero refer to the same
// memory location.


// New takes a type as an argument, allocates enough
// memory to fit a value of that type and returns a
// pointer to it.
// Note: Garbage collected
func one(xPtr *int) {
  *xPtr = 1
}
func main() {
  xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr) // x is 1
}