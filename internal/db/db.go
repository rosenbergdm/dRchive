package db

import (
  "database/sql"
  "fmt"
  "os"
  _ "github.com/mattn/go-sqlite3"
)

func CreateDb(fname string) bool {
  os.Remove(fname)

  fh, err:= os.Create(fname)
  if err != nil {
    fmt.Println("Error creating databse '", fname, "', aborting.")
    return false
  }

  fh.Close()

  db, _ := sql.Open("sqlite3", fname)
  defer db.Close()
  stmt, err := db.Prepare(`
    CREATE TABLE files (
      filepath TEXT PRIMARY KEY ASC,
      mtime INTEGER,
      lastactive
      INTEGER,
      hash TEXT,
      VERSION INTEGER DEFAULT 1
    );`)
  if err != nil {
    fmt.Println("Error writing schema")
    return false
  } else {
    stmt.Exec()
    fmt.Println("Schema written")
  }
  return true
}





