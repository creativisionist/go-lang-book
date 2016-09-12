package main

import (
  "fmt"
  "os"
  "path/filepath"
)

// function passed to Walk is called for every file and folder in root folder (in this case . )
func main() {
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
    })
}
