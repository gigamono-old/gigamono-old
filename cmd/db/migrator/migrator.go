package main

import (
	"flag"

	"github.com/sageflow/sageflow/internal/db/migrations"
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/logs"
)

func main() {
	var up, down bool
	var upTo, downTo string
	var appKind string

	// Set args.
	flag.BoolVar(&up, "up", false, "Migrate the DB to the most recent version available\n")
	flag.BoolVar(&down, "down", false, "Roll back the version by 1\n")
	flag.StringVar(&upTo, "up-to", "", "Migrate the DB to a specific VERSION\n")
	flag.StringVar(&downTo, "down-to", "", "Roll back to a specific VERSION\n")
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

	// Create a migrator.
	migrator := migrations.NewMigrator(app.DB.DB)
	migrator.PrepareMigrations(appKind)

	// log.Println("Last migration =", migrations.GetLastMigration(&app.DB))

	if up {
		if err := migrator.Up(); err != nil {
			logs.FmtPrintln("unsuccessfully migrated:", err)
		}
	} else if down {
		if err := migrator.Down(); err != nil {
			logs.FmtPrintln("unable to roll back:", err)
		}
	} else if upTo != "" {
		if err := migrator.UpTo(upTo); err != nil {
			logs.FmtPrintln("unsuccessfully migrated:", err)
		}
	} else if downTo != "" {
		if err := migrator.DownTo(downTo); err != nil {
			logs.FmtPrintln("unable to roll back:", err)
		}
	} else {
		flag.Usage()
	}
}
