package main

import (
  "fmt"
  "github.com/rosenbergdm/dRchive/internal/db"
)

func main() {
  fmt.Println("Hello, world")
  var res = db.CreateDb("testdb.db")
  fmt.Println("Result ", res, "!")
}

