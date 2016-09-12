package main

import (
  "fmt"
  "container/list"
)

// zero value for List is empty list
// (a *List can also be created using list.New)
// Values are appended to list using PushBack.
// Loop over each item in list by getting first element,
// and following all links until we reach nil.
func main() {
  var x list.List
  x.PushBack(1)
  x.PushBack(2)
  x.PushBack(3)

  for e := x.Front(); e != nil; e=e.Next() {
    fmt.Println(e.Value.(int))
  }
}
