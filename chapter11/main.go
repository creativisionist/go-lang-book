package main

import "fmt"
import  "github.com/golang-book/chapter11/math"
import "github.com/d4l3k/go-pry/pry"

func main() {
  xs := []float64{1,2,3,4}
  avg := math.Average(xs)
  pry.Pry()
  fmt.Println(avg)
}