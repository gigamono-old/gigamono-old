package main

import (
	"flag"

	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/internal/db/seeds"
	"github.com/sageflow/sageflow/pkg/envs"
	"github.com/sageflow/sageflow/pkg/logs"
)

func main() {
	var addAll, removeAll bool
	var seedToAdd, seedToRemove string

	// Set args.
	flag.BoolVar(&addAll, "add-all", false, "Migrate the DB to the most recent version available\n")
	flag.BoolVar(&removeAll, "remove-all", false, "Removes all seeds in table\n")
	flag.StringVar(&seedToAdd, "add", "", "Removes all seeds in table\n")
	flag.StringVar(&seedToRemove, "remove", "", "Removes all seeds in table\n")
	flag.Parse()

	// Set up log status file and load .env file.
	logs.SetStatusLogFile()
	envs.LoadEnvFile()

	// Connect to database.
	db := database.Connect()

	if addAll {
		if err := seeds.AddAll(db); err != nil {
			logs.FmtPrintln("Unable to add all seeds:", err)
		} else {
			logs.FmtPrintln("Successfully added all seeds")
		}
	} else if removeAll {
		if err := seeds.RemoveAll(db); err != nil {
			logs.FmtPrintln("Unable to remove all seeds:", err)
		} else {
			logs.FmtPrintln("Successfully removed all seeds")
		}
	} else if seedToAdd != "" {
		if err := seeds.Add(db, seedToAdd); err != nil {
			logs.FmtPrintf("Unable to add \"%v\" seeds: %v\n", seedToAdd, err)
		} else {
			logs.FmtPrintf("Successfully added \"%v\" seeds\n", seedToAdd)
		}
	} else if seedToRemove != "" {
		if err := seeds.Remove(db, seedToRemove); err != nil {
			logs.FmtPrintf("Unable to remove \"%v\" seeds: %v\n", seedToRemove, err)
		} else {
			logs.FmtPrintf("Successfully removed \"%v\" seeds\n", seedToRemove)
		}
	} else {
		flag.Usage()
	}
}
