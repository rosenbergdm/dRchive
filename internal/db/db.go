package db

import (
  "database/sql"
  "fmt"
  "errors"
  "os"
  "time"
  "path/filepath"
  _ "github.com/mattn/go-sqlite3"
)

type DbEntry struct {
  filepath string
  mtime int64
  lastactive int64
  hash string
}

func NewEntry(db *sql.DB, filepath string, mtime int64, lastactive int64, hash string) error {
  return nil
}

func AddEntry(db *sql.DB, entry *DbEntry) error {
  return nil
}

func UpdateEntry(db *sql.DB, filepath string, mtime int64, lastactive int64, hash string) error {
  return nil
}

func RemoveEntry(db *sql.DB, filepath string) error {
  return nil
}

func GetEntry(db *sql.DB, filepath string) (*DbEntry, error) {
  rows, err := db.Query("Select * from files where filepath = ?", filepath)
  if err != nil {
    return nil, err
  }
  var fpath string
  var mtime time.Time
  var lastactive time.Time
  var hash string
  var version int64

  rows.Next()
  err = rows.Scan(&fpath, &mtime, &lastactive, &hash, &version)
  if err != nil { return nil, err }
  entry := DbEntry{filepath: fpath, mtime: mtime.Unix(), lastactive: lastactive.Unix(), hash: hash}
  if rows.Next() {
    return nil, errors.New("More than 1 row returned")
  }
  return &entry, nil
}

func CreateDb(fname string) (*sql.DB, error) {
  _, err := os.Stat(fname)
  if err == nil {
    fmt.Println("File already exists!")
    return nil, errors.New("File already exists")
  } else if os.IsNotExist(err) {
    dir := filepath.Dir(fname)
    info2, err2 := os.Stat(dir)
    if err2 == nil {
      if !info2.IsDir() {
        fmt.Println("Directory '", dir, "' does not exist")
        return nil, errors.New("Directory does not exist")
      }
      return createDb(fname)
    } else {
      fmt.Println("Error reading '", dir, "'")
      return nil, err2
    }
  } else {
    fmt.Println("Error reading '", fname, "'")
    return nil, errors.New("Error reading file")
  }
}

func createDb(fname string) (*sql.DB, error) {
  fh, err:= os.Create(fname)
  if err != nil {
    fmt.Println("Error creating database '", fname, "', aborting.")
    return nil, err
  }
  fh.Close()
  db, err := sql.Open("sqlite3", fname)
  if err != nil {
    fmt.Println("Error opening database '", fname, "'")
    return nil, err
  }
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
    return nil, err
  } else {
    stmt.Exec()
    fmt.Println("Schema written")
  }
  return db, nil
}

func OpenDb(fname string) (*sql.DB, error) {
  db, err := sql.Open("sqlite3", fname)
  if err != nil {
    fmt.Println("Error opening database '", fname, "'")
    return nil, err
  }
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
    return nil, err
  } else {
    stmt.Exec()
    fmt.Println("Schema written")
  }
  return db, nil
}







