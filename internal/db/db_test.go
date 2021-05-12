package db

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"
)

func TestNewEntry(t *testing.T) {
	dbfile, err := os.CreateTemp("", "tempdb.*.db")
	if err != nil {
		t.Fatal("Could not create a tempfile")
	}
  dbfile.Close()
	// defer os.Remove(dbfile.Name())
	os.Remove(dbfile.Name())
	db, err := CreateDb(dbfile.Name())
	if err != nil {
		t.Fatalf(fmt.Sprintf("Could not write test db '%s'", dbfile.Name()))
	}
	res := db.NewEntry("/bin/bash", time.Now(), time.Now(), "a4221a3a4344e4f86e70d1e475e7ccee")
	if res != nil {
		t.Fatalf("Could not create entry")
	}
  time.Sleep(time.Duration(100000))
  fmt.Printf("/usr/bin/sqlite3 %s -cmd 'SELECT * from files' -cmd '.exit 1'", dbfile.Name())
	out, err := exec.Command(fmt.Sprintf("/usr/bin/sqlite3 %s -cmd 'SELECT * from files' -cmd '.exit 1'", dbfile.Name())).Output()
  fmt.Println(out)
	expected := regexp.MustCompile(`^/bin/bash|\d+|\d+|/.bashrc|0|0|484e75d8a73cd9d782f080c028af50fa`)
	if !expected.Match(out) {
		t.Fatalf(fmt.Sprintf("Database result didn't match '%s'", out))
	} else if err != nil {
		t.Fatalf("Error checking output")
	}
}
