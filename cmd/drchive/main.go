package main

import (
  "fmt"
  "os"
  "github.com/rosenbergdm/dRchive/internal/db"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("USAGE: ", os.Args[0], " <DATABASE>")
    os.Exit(1)
  }
  dbname := os.Args[1]
  db, err := db.CreateDb(dbname)
  if err != nil {
    fmt.Println("Error creating DB!  Aborting")
    os.Exit(99)
  }
  db.Close()
  fmt.Println("Success")
}

