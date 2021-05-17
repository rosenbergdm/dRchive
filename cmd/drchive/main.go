package main

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `Usage:
  drchive [--debug] --config=<configdir> --source=<srcdir> --target=<targetdir> --image=<imagedir>
  drchive [--debug] --configfile=<configfile>
  drchive [-h | --help]
  drchive

Run the dRchive program

Options:
  -h --help                   Display usage
  --debug                     Set maximal output
  -v --verbose                Increase verbosity
  --config=<configfile>       Directory containing configuration components
  --source=<srcdir>           Directory to read from
  --target=<targetdir>        Directory to archive into
  --image=<imagedir>          Image directory
  --configfile=<configfile>   File containing all configuration params`

	opts, err := docopt.ParseArgs(usage, os.Args[1:], "0.0.1")

	if err == nil {
		fmt.Println("Arguments parsed successfully")
		fmt.Println(opts)
	} else {
		fmt.Printf("%+v\n", err)
		fmt.Println("Arguments NOT parsed successfully")
		fmt.Println(err.Error())
	}

	os.Exit(0)
	// 	if len(os.Args) < 2 {
	// 		fmt.Println("USAGE: ", os.Args[0], " <DATABASE>")
	// 		os.Exit(1)
	// 	}
	// 	dbname := os.Args[1]
	// 	db, err := db.CreateDb(dbname)
	// 	if err != nil {
	// 		fmt.Println("Error creating DB!  Aborting")
	// 		os.Exit(99)
	// 	}
	// 	db.Close()
	// 	fmt.Println("Success")
}
