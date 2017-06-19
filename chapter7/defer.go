package main

import "fmt"

func first() {
  fmt.Println("1st")
}
func second() {
  fmt.Println("2nd")
}
func main() {
  defer second()
  first()
}

// Often used when resources need to be freed in some way
// e.g. Open a file, need to make sure to close it later
// f, _ := os.Open(filename)
// defer f.Close()

// This has 3 advantages: (1) it keeps our Close call near
// our Open call so its easier to understand, (2) if our function
// had multiple return statements (perhaps one in
// an if and one in an else) Close will happen before
// both of them and (3) deferred functions are run even if
// a run-time panic occurs.