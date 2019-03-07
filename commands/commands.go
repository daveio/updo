package commands

import (
	"fmt"
	"github.com/daveio/updo/storage"
	"github.com/daveio/updo/verbose"
)

var (
	V = verbose.V
)

func Run(conf storage.Conf, dryRunFlag bool) (error error) {
	// TODO commands::Run
	_, _ = conf, dryRunFlag
	return nil
}

func Hello() {
	fmt.Println("https://github.com/daveio/updo")
}