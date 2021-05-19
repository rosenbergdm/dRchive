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
	log.ConfigLogger(log.InfoLevel, os.Stdout)
}

func (db *FileDb) NewEntry(filepath string, mtime time.Time, lastactive time.Time, hash string) error {
	_, err := db.Exec("INSERT INTO files (filepath, mtime, lastactive, hash) VALUES (?, ?, ?, ?)", filepath, mtime.Unix(), lastactive.Unix(), hash)
	if err != nil {
		return err
	}
	log.LogInfo(log.Fields{"filepath": filepath}, "New Insertion")
	return nil
}

func AddEntry(db *FileDb, entry *DbEntry) error {
	_, err := db.Exec("INSERT INTO files (filepath, mtime, lastactive, hash) VALUES (?, ?, ?, ?)", entry.filepath, entry.mtime, entry.lastactive, entry.hash)
	if err != nil {
		return err
	}
	log.LogInfo(log.Fields{"filepath": entry.filepath}, "New insertion")
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
		log.LogWarn(fields, "Cannot update entry without any new fields")
		return err
	}
	log.LogInfo(fields, "Record updated")
	return nil
}

func RemoveEntry(db *FileDb, filepath string) error {
	_, err := db.Exec("DELETE FROM files WHERE filepath=?", filepath)
	if err != nil {
		log.LogWarn(log.Fields{"filepath": filepath, "error": err}, "Unable to delete")
		return err
	}
	log.LogInfo(log.Fields{"filepath": filepath}, "Deletion successful")
	return nil
}

func GetEntry(db *FileDb, filepath string) (*DbEntry, error) {
	rows, err := db.Query("Select * from files where filepath = ?", filepath)
	if err != nil {
		log.LogFatal(log.Fields{"filepath": filepath}, "Unable to query database")
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
		log.LogWarn(log.Fields{"filepath": filepath}, "No entry matches query")
		return nil, err
	}
	entry := DbEntry{filepath: fpath, mtime: mtime.Unix(), lastactive: lastactive.Unix(), hash: hash}
	if rows.Next() {
		log.LogFatal(log.Fields{"filepath": filepath}, "Multiple entries with identical paths found!")
	}
	log.LogInfo(log.Fields{"filepath": filepath}, "Found entry")
	return &entry, nil
}

func CreateDb(fname string) (*FileDb, error) {
	_, err := os.Stat(fname)
	if err == nil {
		log.LogFatal(log.Fields{"dbfile": fname}, "File already exists!")
	} else if os.IsNotExist(err) {
		dir := filepath.Dir(fname)
		info2, err2 := os.Stat(dir)
		if err2 == nil {
			if !info2.IsDir() {
				log.LogFatal(log.Fields{"directory": dir}, "Directory does not exist")
			}
			return createDb(fname)
		} else {
			log.LogFatal(log.Fields{"directory": dir}, "Cannot read directory")
		}
	} else {
		log.LogFatal(log.Fields{"directory": fname}, "Cannot read directory")
	}
	return nil, errors.New("Unknown error")
}

func createDb(fname string) (*FileDb, error) {
	fh, err := os.Create(fname)
	if err != nil {
		log.LogFatal(log.Fields{"fname": fname}, "Error creating file")
	}
	fh.Close()
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		log.LogFatal(log.Fields{"fname": fname}, "Error opening file with sqlite3")
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
		log.LogFatal(log.Fields{"fname": fname, "error": err}, "Error writing schema")
	} else {
		log.LogInfo(log.Fields{"dbfile": fname}, "Schema written")
	}
	fdb := &FileDb{db}
	return fdb, nil
}

func OpenDb(fname string) (*FileDb, error) {
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		log.LogFatal(log.Fields{"fname": fname}, "Error opening file with sqlite3")
	}
	log.LogInfo(log.Fields{"dbfile": fname}, "DB opened")
	fdb := &FileDb{db}
	return fdb, nil
}
