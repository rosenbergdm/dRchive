package db

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/rosenbergdm/dRchive/internal/log"

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

func init() {
	log.Config(log.InfoLevel, os.Stdout)
}

func (db *FileDb) NewEntry(filepath string, mtime time.Time, lastactive time.Time, hash string) error {
	_, err := db.Exec("INSERT INTO files (filepath, mtime, lastactive, hash) VALUES (?, ?, ?, ?)", filepath, mtime.Unix(), lastactive.Unix(), hash)
	if err != nil {
		return err
	}
	log.Info("New Insertion", log.Fields{"filepath": filepath})
	return nil
}

func AddEntry(db *FileDb, entry *DbEntry) error {
	_, err := db.Exec("INSERT INTO files (filepath, mtime, lastactive, hash) VALUES (?, ?, ?, ?)", entry.filepath, entry.mtime, entry.lastactive, entry.hash)
	if err != nil {
		return err
	}
	log.Info("New insertion", log.Fields{"filepath": entry.filepath})
	return nil
}

func UpdateEntry(db *FileDb, filepath string, mtime time.Time, lastactive time.Time, hash string) error {
	var err error
	fields := map[string]interface{}{"filepath": filepath}
	if mtime.IsZero() {
		if lastactive.IsZero() {
			if hash != "" {
				_, err = db.Exec("UPDATE files set hash=? where filepath=?", hash, filepath)
				fields["hash"] = hash
			} else {
				err = errors.New("No fields to update")
			}
		} else {
			fields["lastactive"] = lastactive.String()
			if hash != "" {
				fields["hash"] = hash
				_, err = db.Exec("UPDATE files set hash=?, lastactive=? where filepath=?", hash, lastactive.Unix(), filepath)
			} else {
				_, err = db.Exec("UPDATE files SET lastactive=? WHERE filepath=?", lastactive.Unix(), filepath)
			}
		}
	} else {
		fields["mtime"] = mtime.String()
		if lastactive.IsZero() {
			if hash != "" {
				fields["hash"] = hash
				_, err = db.Exec("UPDATE files set hash=?, mtime=? where filepath=?", hash, mtime.Unix(), filepath)
			} else {
				_, err = db.Exec("Update files set mtime=? where filepath=?", mtime.Unix(), filepath)
			}
		} else {
			fields["lastactive"] = lastactive.String()
			if hash != "" {
				fields["hash"] = hash
				_, err = db.Exec("UPDATE files set hash=?, lastactive=?, mtime=? where filepath=?", hash, lastactive.Unix(), mtime.Unix(), filepath)
			} else {
				_, err = db.Exec("UPDATE files SET lastactive=?, mtime=? WHERE filepath=?", lastactive.Unix(), mtime.Unix(), filepath)
			}
		}
	}
	if err != nil {
		fields["errror"] = err
		log.Warn("Cannot update entry without any new fields", fields)
		return err
	}
	log.Info("Record updated", fields)
	return nil
}

func RemoveEntry(db *FileDb, filepath string) error {
	_, err := db.Exec("DELETE FROM files WHERE filepath=?", filepath)
	if err != nil {
		log.Warn("Unable to delete", log.Fields{"filepath": filepath, "error": err})
		return err
	}
	log.Info("Deletion successful", log.Fields{"filepath": filepath})
	return nil
}

func GetEntry(db *FileDb, filepath string) (*DbEntry, error) {
	rows, err := db.Query("Select * from files where filepath = ?", filepath)
	if err != nil {
		log.Fatal("Unable to query database", log.Fields{"filepath": filepath})
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
		log.Warn("No entry matches query", log.Fields{"filepath": filepath})
		return nil, err
	}
	entry := DbEntry{filepath: fpath, mtime: mtime.Unix(), lastactive: lastactive.Unix(), hash: hash}
	if rows.Next() {
		log.Fatal("Multiple entries with identical paths found!", log.Fields{"filepath": filepath})
	}
	log.Info("Found entry", log.Fields{"filepath": filepath})
	return &entry, nil
}

func CreateDb(fname string) (*FileDb, error) {
	_, err := os.Stat(fname)
	if err == nil {
		log.Fatal("File already exists!", log.Fields{"dbfile": fname})
	} else if os.IsNotExist(err) {
		dir := filepath.Dir(fname)
		info2, err2 := os.Stat(dir)
		if err2 == nil {
			if !info2.IsDir() {
				log.Fatal("Directory does not exist", log.Fields{"directory": dir})
			}
			return createDb(fname)
		} else {
			log.Fatal("Cannot read directory", log.Fields{"directory": dir})
		}
	} else {
		log.Fatal("Cannot read directory", log.Fields{"directory": fname})
	}
	return nil, errors.New("Unknown error")
}

func createDb(fname string) (*FileDb, error) {
	fh, err := os.Create(fname)
	if err != nil {
		log.Fatal("Error creating file", log.Fields{"fname": fname})
	}
	fh.Close()
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		log.Fatal("Error opening file with sqlite3", log.Fields{"fname": fname})
	}
	stmt, err := db.Prepare(`
    CREATE TABLE files (
      filepath TEXT PRIMARY KEY ASC,
      mtime INTEGER,
      lastactive INTEGER,
      hash TEXT,
      VERSION INTEGER DEFAULT 1
    );`)
	stmt.Exec()
	if err != nil {
		log.Fatal("Error writing schema", log.Fields{"fname": fname, "error": err})
	} else {
		log.Info("Schema written", log.Fields{"dbfile": fname})
	}
	fdb := &FileDb{db}
	return fdb, nil
}

func OpenDb(fname string) (*FileDb, error) {
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		log.Fatal("Error opening file with sqlite3", log.Fields{"fname": fname})
	}
	log.Info("DB opened", log.Fields{"dbfile": fname})
	fdb := &FileDb{db}
	return fdb, nil
}
