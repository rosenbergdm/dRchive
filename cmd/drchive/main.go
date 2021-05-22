package main

import (
	"os"

	docopt "github.com/docopt/docopt-go"

	"github.com/rosenbergdm/drchive/internal/log"
)

func init() {
	log.Config(log.WarnLevel, os.Stdout)
}

func main() {
	usage := `Usage:
  drchive [-vhdq] [--debug] --config=<configdir> --source=<srcdir> --target=<targetdir> --image=<imagedir>
  drchive [-vhdq] [--debug] --configfile=<configfile>
  drchive [-h | --help]
  drchive

Run the drchive program

Options:
  --help -h                   Display usage
  --debug                     Set maximal output
  --verbose                   Increase verbosity
  --quiet -q                  Suppress output
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
	// log.LogError(logrus.Fields{}, "help************************")
	// log.LogWarn(logrus.Fields{"testing": "123"}, "help")
	opts, err := docopt.ParseArgs(usage, os.Args[1:], "0.0.1")
	if err != nil {
		log.Fatal("Unable to parse commandline args", log.Fields{})
	}

	if opts["--debug"].(bool) == true {
		log.Config(log.InfoLevel, os.Stdout)
		log.Info("Sending DEBUG output to STDOUT", log.Fields{})
	}

	log.Info("Arguments parsed successfully", log.Fields{"options": opts})

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
