package file

import (
	"errors"
	"path/filepath"

	"github.com/rosenbergdm/dRchive/internal/db"
	"github.com/rosenbergdm/dRchive/internal/log"
)

type SymLink struct {
	source string
	Target string
}

type Runner struct {
	srcRoot string
	tgtRoot string
	links   []SymLink
	db      *db.FileDb
	errs    []error
}

// func NewRunner(srcRoot, tgtRoot, dbFile) (*FileDb, error) {

// }

// func OpenRunner() {}

type fake struct {
	e error
	d *db.FileDb
}

func Newfake() fake {
	log.Info("", nil)
	fn, _ := filepath.Abs(".")
	db, _ := db.OpenDb(fn)
	f := fake{e: errors.New(""), d: db}
	return f
}
