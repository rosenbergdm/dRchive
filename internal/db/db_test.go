package db

import (
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"

	"github.com/rosenbergdm/dRchive/internal/log"
)

func setupTestDb() (*FileDb, func(), string) {
	log.Config(log.ErrorLevel, os.Stdout)
	dbfile, err := os.CreateTemp("", "tempdb.*.db")
	if err != nil {
		log.Fatal("Could not create a tempfile", nil)
	}
	dbfile.Close()
	os.Remove(dbfile.Name())
	db, err := CreateDb(dbfile.Name())
	if err != nil {
		log.Fatal("Could not write test db", log.Fields{"dbfile": dbfile.Name(), "error": err.Error()})
	}
	tearDown := func() {
		db.Close()
		os.Remove(dbfile.Name())
	}
	log.Config(log.PanicLevel, os.Stdout)
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
	}
	rows, err2 := db.Query("SELECT COUNT * FROM files")
	if err2 != nil {
		t.Fatalf("error querying test db '%s'", err.Error())
	}
	var count int64
	rows.Scan(&count)
	if count != 1 {
		t.Fatal("Multiple returns for a search!")
	}
}
