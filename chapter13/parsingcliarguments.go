package main

import (
  "fmt"
  "flag"
  "math/rand"
)

func main() {
  // Define flags
  maxp := flag.Int("max", 6, "the max value")
  // Parse
  flag.Parse()
  // Generate # b/w 0 and max
  fmt.Println(rand.Intn(*maxp))
}


// go run parsingcliarguments.go -max=100
