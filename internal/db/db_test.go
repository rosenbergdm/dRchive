package db

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"

	"github.com/rosenbergdm/dRchive/internal/log"
)

func countEntries(db *FileDb, t *testing.T) int64 {
	row := db.QueryRow("SELECT COUNT(*) FROM files")
	var count int64
	err := row.Scan(&count)
	if err != nil {
		t.Fatalf("error querying test db '%v'", err)
	}
	return count
}

func setupTestDb() (*FileDb, func(), string) {
	debugTests := os.Getenv("DEBUG_TESTS")
	saveDBs := os.Getenv("DEBUG_SAVE_DBS")
	if debugTests != "" {
		log.Config(log.InfoLevel, os.Stdout)
	} else {
		log.Config(log.ErrorLevel, os.Stdout)
	}
	dbfile, err := os.CreateTemp("", "tempdb.*.db")
	if err != nil {
		log.Fatal("setupTestDb error: Could not create a tempfile", nil)
	}
	dbfile.Close()
	os.Remove(dbfile.Name())
	db, err := CreateDb(dbfile.Name())
	if err != nil {
		log.Fatal("setupTestDb error: Could not write test db", log.Fields{"dbfile": dbfile.Name(), "error": err.Error()})
	}
	tearDown := func() {
		db.Close()
		if debugTests != "" {
			log.Config(log.InfoLevel, os.Stdout)
		} else {
			log.Config(log.PanicLevel, os.Stdout)
		}
		if saveDBs != "" {
			fmt.Printf("File is '%s'\n\n", dbfile.Name())
		} else {
			os.Remove(dbfile.Name())
		}
	}
	return db, tearDown, dbfile.Name()
}

func TestNewEntry(t *testing.T) {
	var db *FileDb
	db, tearDown, dbfile := setupTestDb()
	defer tearDown()
	res := db.NewEntry("/bin/bash", time.Now(), time.Now(), "a4221a3a4344e4f86e70d1e475e7ccee")
	if res != nil {
		t.Fatal("Could not create entry")
	}
	out, err := exec.Command("/usr/bin/sqlite3", dbfile, "-cmd", "SELECT * from files", "-cmd", ".exit 0").Output()
	outText := string(out)
	expected := regexp.MustCompile(`^/bin/bash\|\d+\|\d+\|a4221a3a4344e4f86e70d1e475e7ccee\|1$`)
	if expected.MatchString(outText) {
		t.Fatalf("Database result didn't match\nOUTPUT  : %s\nEXPECTED: %s\n", outText, `^/bin/bash\|\d+\|\d+\|a4221a3a4344e4f86e70d1e475e7ccee\|1$`)
	} else if err != nil {
		t.Fatalf("Error checking output: %s", err)
	}
}

func TestNewEntryNoDups(t *testing.T) {
	db, tearDown, _ := setupTestDb()
	defer tearDown()
	err := db.NewEntry("/bin/bash", time.Time{}, time.Time{}, "a4221a3a4344e4f86e70d1e475e7ccee")
	if err != nil {
		t.Fatalf("Broken Test: %s", err.Error())
	}
	err = db.NewEntry("/bin/bash", time.Time{}, time.Time{}, "a4221a3a4344e4f86e70d1e475e7ccff")
	if err == nil {
		t.Fatal("Allowed a duplicate entries")
		os.Exit(99)
	}
	row := db.QueryRow("SELECT COUNT(*) FROM files")
	var count int64
	err = row.Scan(&count)
	if err != nil {
		t.Fatalf("error querying test db '%v'", err)
	}
	if count != 1 {
		t.Fatalf("Multiple returns for a search: count=%v", count)
	}
}

func TestAddEntry(t *testing.T) {
	entries := []*DbEntry{
		{filepath: "/bin/bash", mtime: 0, lastactive: 0, hash: "a4221a3a4344e4f86e70d1e475e7ccee"},
		{filepath: "/bin/zsh", mtime: 0, lastactive: 0, hash: "99f587cd90ebb1d8ae04ea45bcc66481"},
	}

	db, tearDown, _ := setupTestDb()
	defer tearDown()
	for i := range entries {
		db.AddEntry(entries[i])
	}
	count := countEntries(db, t)
	if count != 2 {
		t.Fatalf("Wrong number of entries added, should be 2 but is %d", count)
	}
}

func TestRemoveEntry(t *testing.T) {

}

func TestGetEntry(t *testing.T) {

}

func TestOpenDb(t *testing.T) {

}
