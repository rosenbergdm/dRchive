package main

import (
	// "flag"
	"fmt"
	"os"

	"github.com/rosenbergdm/dRchive/internal/log"

	docopt "github.com/docopt/docopt-go"
	"github.com/sirupsen/logrus"
)

func main() {
	usage := `Usage:
  drchive [--debug] --config=<configdir> --source=<srcdir> --target=<targetdir> --image=<imagedir>
  drchive [--debug] --configfile=<configfile>
  drchive [-h | --help]
  drchive

Run the dRchive program

Options:
  --help                      Display usage
  --debug                     Set maximal output
  --verbose                   Increase verbosity
  --config=<configdir>        Directory containing configuration components
  --source=<srcdir>           Directory to read from
  --target=<targetdir>        Directory to archive into
  --image=<imagedir>          Image directory
  --configfile=<configfile>   File containing all configuration params
  --version                   Version
`
	// help := flag.Bool("help", false, "Display usage")
	// debug := flag.Bool("debug", false, "Set maximal output")
	// verbose := flag.Bool("verbose", false, "Set verbose output")
	// version := flag.Bool("version", false, "Display version number")
	// configfile := flag.String("configfile", "", "Configuration file")

	// flag.Parse()
	// fmt.Println("'help' has value: ", *help)
	// fmt.Println("'debug' has value: ", *debug)
	// fmt.Println("'verbose' has value: ", *verbose)
	// fmt.Println("'version' has value: ", *version)
	// fmt.Println("'configfile' has value: ", *configfile)

	// fmt.Println(usage)

	// fmt.Print("\n---------------------\n\n")
	log.LogError(logrus.Fields{}, "help************************")
	log.LogWarn(logrus.Fields{"testing": "123"}, "help")
	opts, err := docopt.ParseArgs(usage, os.Args[1:], "0.0.1")

	if err == nil {
		fmt.Println("Arguments parsed successfully")
		docopt.PrintHelpOnly(err, usage)
		fmt.Println(opts)
	} else {
		fmt.Printf("%+v\n", err)
		fmt.Println("Arguments NOT parsed successfully")
		fmt.Println(err.Error())
	}
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
