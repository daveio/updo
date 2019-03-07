package main

import (
	"fmt"
	"github.com/daveio/updo/commands"
	"github.com/daveio/updo/storage"
	"github.com/daveio/updo/verbose"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	V                  = verbose.V
	appVersion         = fmt.Sprintf("%d.%d.%d", AppVersionMajor, AppVersionMinor, AppVersionPatch)
	appVersionWithDate = fmt.Sprintf("%s %d", appVersion, AppVersionDate)
	app                = kingpin.
				New("updo", "Update all the updatables.")
	mVerbose = app.
			Flag("verbose", "Show more detail.").
			Short('v').
			Bool()
	runCommand = app.
			Command("run", "Short form: 'r'. Run the update process.").
			Alias("r")
	dryRunFlag = runCommand.
			Flag("dry-run", "Print the commands which would be executed, but don't actually run them").
			Short('d').
			Bool()
	helloCommand = app.
			Command("hello", "Used by shell plugins like zsh-updo to identify this binary. Prints GitHub URL.")
)

func init() {
	storage.InitStorage("updo")
}

func main() {
	app.Version(appVersionWithDate)
	appCmd, appErr := app.Parse(os.Args[1:])
	verbose.InitV(*mVerbose, appVersion)
	V("Verbose mode enabled")
	conf, errL := storage.LoadConfig()
	if errL != nil {
		panic(errL)
	}
	if len(conf.Apps) < 1 {
		conf.Apps = make(map[string]storage.App)
	}
	errBuiltins := storage.LoadBuiltinApps(&conf)
	if errBuiltins != nil {
		panic(errBuiltins)
	}
	switch kingpin.MustParse(appCmd, appErr) {
	case runCommand.FullCommand():
		V("starting the update process")
		if *dryRunFlag {
			V("dry run: not executing anything")
		} else {
			V("executing update commands")
		}
		errRun := commands.Run(conf, *dryRunFlag)
		if errRun != nil {
			panic(errRun)
		}
		V("finished update process")
	case helloCommand.FullCommand():
		commands.Hello()
	default:
		V("BUG: invalid command uncaught by CLI parser")
	}
	errSa := storage.SaveConfig(&conf)
	if errSa != nil {
		panic(errSa)
	}
}
