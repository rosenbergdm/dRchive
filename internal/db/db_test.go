package db

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"
)

func setupTestDb() (*FileDb, func(), string) {
	dbfile, err := os.CreateTemp("", "tempdb.*.db")
	if err != nil {
		log.Fatal("Could not create a tempfile")
	}
	dbfile.Close()
	os.Remove(dbfile.Name())
	db, err := CreateDb(dbfile.Name())
	if err != nil {
		log.Fatalf("Could not write test db '%s': '%s'", dbfile.Name(), err.Error())
	}
	tearDown := func() {
		db.Close()
		os.Remove(dbfile.Name())
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
