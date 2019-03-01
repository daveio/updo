package main

import (
	"fmt"
	"github.com/tucnak/store"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"time"
)

var (
	app = kingpin.
		New("updo", "Update all the things!")
	debug = app.
		Flag("debug", "Enable debug mode.").
		Short('d').
		Bool()
	updateCommand = app.
			Command("update", "Update everything except updo itself.").
			Alias("u")
	selfUpdateCommand = app.
			Command("self-update", "Update updo itself (and nothing else)").
			Alias("g")
)

type Directive struct {
	Key string `json:"key"`
	Value  string `json:"value"`
}

type AppConfig struct {
	Directives []Directive `json:"directive"`
}

func init() {
	store.Init("updo")
}

func loadConfig() AppConfig {
	var appConfig AppConfig
	err := store.Load("config.json", &appConfig)
	if err != nil {
		panic(err)
	}
	return appConfig
}

func saveConfig(appConfig *AppConfig) {
	err := store.Save("config.json", &appConfig)
	if err != nil {
		panic(err)
	}
}

func main() {
	appConfig := loadConfig()
	if *debug {
		fmt.Println("Debug mode enabled.")
	}
	app.Version("0.0.1")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case updateCommand.FullCommand():
		fmt.Printf("UPDATING ALL THE THINGS")
		fmt.Println()
	case selfUpdateCommand.FullCommand():
		fmt.Printf("UPDATING UPDO ITSELF")
		fmt.Println()
	}
	time.Sleep(1 * time.Second)
	saveConfig(&appConfig)
}
