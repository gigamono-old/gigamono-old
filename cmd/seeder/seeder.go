package main

import (
	"flag"

	"github.com/gigamono/gigamono/internal/db/seeds"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"
)

func main() {
	var addAll, removeAll bool
	var seedToAdd, seedToRemove, appKind string

	// Set args.
	flag.BoolVar(&addAll, "add-all", false, "Migrate the DB to the most recent version available\n")
	flag.BoolVar(&removeAll, "remove-all", false, "Removes all seeds in table\n")
	flag.StringVar(&seedToAdd, "add", "", "Removes all seeds in table\n")
	flag.StringVar(&seedToRemove, "remove", "", "Removes all seeds in table\n")
	flag.StringVar(&appKind, "type", "", "Specify the application kind (resource or auth)\n")
	flag.StringVar(&appKind, "t", "", "Specify the application kind (resource or auth)\n")
	flag.Parse()

	// Set default app kind.
	if appKind == "" {
		appKind = "Resource"
	}

	// Initialise app.
	app, err := inits.NewApp(appKind)
	if err != nil {
		logs.FmtPrintln(err)
		return
	}

	// Create a seeder.
	seeder := seeds.NewSeeder(&app.DB, appKind)

	if addAll {
		if err := seeder.AddAll(); err != nil {
			logs.FmtPrintln("unable to add all seeds:", err)
		} else {
			logs.FmtPrintln("successfully added all seeds")
		}
	} else if removeAll {
		if err := seeder.RemoveAll(); err != nil {
			logs.FmtPrintln("unable to remove all seeds:", err)
		} else {
			logs.FmtPrintln("successfully removed all seeds")
		}
	} else if seedToAdd != "" {
		if err := seeder.Add(seedToAdd); err != nil {
			logs.FmtPrintf("unable to add \"%v\" seeds: %v\n", seedToAdd, err)
		} else {
			logs.FmtPrintf("successfully added \"%v\" seeds\n", seedToAdd)
		}
	} else if seedToRemove != "" {
		if err := seeder.Remove(seedToRemove); err != nil {
			logs.FmtPrintf("unable to remove \"%v\" seeds: %v\n", seedToRemove, err)
		} else {
			logs.FmtPrintf("successfully removed \"%v\" seeds\n", seedToRemove)
		}
	} else {
		flag.Usage()
	}
}
