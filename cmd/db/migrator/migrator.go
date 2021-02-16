package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/sageflow/sageflow/internal/db/migrations"
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/logs"
)

func main() {
	var up, down bool
	var upTo, downTo int
	var appKind string

	// Set args.
	flag.BoolVar(&up, "up", false, "Migrate the DB to the most recent version available\n")
	flag.BoolVar(&down, "down", false, "Roll back the version by 1\n")
	flag.IntVar(&upTo, "up-to", 0, "Migrate the DB to a specific VERSION\n")
	flag.IntVar(&downTo, "down-to", 0, "Roll back to a specific VERSION\n")
	flag.StringVar(&appKind, "kind", "", "Specify the application kind (resource or auth)\n")
	flag.StringVar(&appKind, "k", "", "Specify the application kind (resource or auth)\n")
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

	// Create a migrator.
	migrator := migrations.NewMigrator(&app.DB, appKind)

	log.Println("Last migration =", migrations.GetLastMigration(&app.DB))

	if up {
		if err := migrator.Migrate(); err != nil {
			logs.FmtPrintln("Unsuccessfully migrated:", err)
		} else {
			logs.FmtPrintln("Successfully migrated to the most recent version")
		}
	} else if down {
		if err := migrator.RollbackLast(); err != nil {
			logs.FmtPrintln("Unable to roll back:", err)
		} else {
			logs.FmtPrintln("Rolled back last migration")
		}
	} else if upTo > 0 {
		if err := migrator.MigrateTo(strconv.Itoa(upTo)); err != nil {
			logs.FmtPrintln("Unsuccessfully migrated:", err)
		} else {
			logs.FmtPrintln("Successfully migrated to version")
		}
	} else if downTo > 0 {
		if err := migrator.RollbackTo(strconv.Itoa(downTo)); err != nil {
			logs.FmtPrintln("Unable to roll back:", err)
		} else {
			logs.FmtPrintln("Rolled back migration to version")
		}
	} else {
		flag.Usage()
	}
}
