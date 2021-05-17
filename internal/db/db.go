package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
  log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3" // For db driver
)

// Each file entry in the database
type DbEntry struct {
	filepath   string
	mtime      int64
	lastactive int64
	hash       string
}

// SQL Database containing the files and hashes
type FileDb struct {
	*sql.DB
}

func (db *FileDb) NewEntry(filepath string, mtime time.Time, lastactive time.Time, hash string) error {
	_, err := db.Exec("INSERT INTO files (filepath, mtime, lastactive, hash) VALUES (?, ?, ?, ?)", filepath, mtime.Unix(), lastactive.Unix(), hash)
	if err != nil {
		return err
	}
  log.WithFields(log.Fields{"Filepath": filepath}).Info("New insertion")
	return nil
}

func AddEntry(db *FileDb, entry *DbEntry) error {
	_, err := db.Exec("INSERT INTO files (filepath, mtime, lastactive, hash) VALUES (?, ?, ?, ?)", entry.filepath, entry.mtime, entry.lastactive, entry.hash)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEntry(db *FileDb, filepath string, mtime time.Time, lastactive time.Time, hash string) error {
	var err error
	if mtime.IsZero() {
		if lastactive.IsZero() {
			if hash == "" {
				_, err = db.Exec("UPDATE files set hash=? where filepath=?", hash, filepath)
			} else {
				err = nil
			}
		} else {
			if hash == "" {
				_, err = db.Exec("UPDATE files set hash=?, lastactive=? where filepath=?", hash, lastactive.Unix(), filepath)
			} else {
				_, err = db.Exec("UPDATE files SET lastactive=? WHERE filepath=?", lastactive.Unix(), filepath)
			}
		}
	} else {
		if lastactive.IsZero() {
			if hash == "" {
				_, err = db.Exec("UPDATE files set hash=?, mtime=?, where filepath=?", hash, mtime.Unix(), filepath)
			} else {
				err = nil
			}
		} else {
			if hash == "" {
				_, err = db.Exec("UPDATE files set hash=?, lastactive=?, mtime=? where filepath=?", hash, lastactive.Unix(), mtime.Unix(), filepath)
			} else {
				_, err = db.Exec("UPDATE files SET lastactive=?, mtime=? WHERE filepath=?", lastactive.Unix(), mtime.Unix(), filepath)
			}
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func RemoveEntry(db *FileDb, filepath string) error {
	_, err := db.Exec("DELETE FROM files WHERE filepath=?", filepath)
	if err != nil {
		return err
	}
	return nil
}

func GetEntry(db *FileDb, filepath string) (*DbEntry, error) {
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
	if err != nil {
		return nil, err
	}
	entry := DbEntry{filepath: fpath, mtime: mtime.Unix(), lastactive: lastactive.Unix(), hash: hash}
	if rows.Next() {
		return nil, errors.New("More than 1 row returned")
	}
	return &entry, nil
}

func CreateDb(fname string) (*FileDb, error) {
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

func createDb(fname string) (*FileDb, error) {
	fh, err := os.Create(fname)
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
	fdb := &FileDb{db}
	return fdb, nil

}

func OpenDb(fname string) (*FileDb, error) {
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
      lastactive INTEGER,
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
	fdb := &FileDb{db}
	return fdb, nil
}
