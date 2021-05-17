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
	os.Remove(dbfile.Name())
	db, err := CreateDb(dbfile.Name())
	if err != nil {
		t.Fatalf(fmt.Sprintf("Could not write test db '%s'", dbfile.Name()))
	}
	res := db.NewEntry("/bin/bash", time.Now(), time.Now(), "a4221a3a4344e4f86e70d1e475e7ccee")
	if res != nil {
		t.Fatalf("Could not create entry")
	}
	out, err := exec.Command("/usr/bin/sqlite3", dbfile.Name(), "-cmd", "SELECT * from files", "-cmd", ".exit 0").Output()
	outText := string(out)
	expected := regexp.MustCompile(`^/bin/bash\|\d+\|\d+\|a4221a3a4344e4f86e70d1e475e7ccee\|1$`)
	if expected.MatchString(outText) {
		t.Fatalf(fmt.Sprintf("Database result didn't match\nOUTPUT  : %s\nEXPECTED: %s\n", outText, `^/bin/bash\|\d+\|\d+\|a4221a3a4344e4f86e70d1e475e7ccee\|1$`))
	} else if err != nil {
		t.Fatalf(fmt.Sprintf("Error checking output: %s", err))
	}
}
