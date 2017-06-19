package main

import (
  "fmt"
  "errors"
)

func main() {
  err := errors.New("own customer error message")
  fmt.Println(err)
}
