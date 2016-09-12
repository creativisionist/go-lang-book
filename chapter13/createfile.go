package main

import (
  "os"
)

func main() {
  file, err := os.Create("holler.txt")
  if err != nil {
    // handle the error here
    return
  }
  defer file.Close()

  file.WriteString("Holler!")
}