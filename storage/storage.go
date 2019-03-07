package storage

import (
	"github.com/daveio/updo/verbose"
	"github.com/tucnak/store"
)

var (
	V = verbose.V
)

func InitStorage(appName string) {
	store.Init(appName)
}

func LoadConfig() (conf Conf, error error) {
	var c Conf
	err := store.Load("updo.json", &c)
	return c, err
}

func SaveConfig(conf *Conf) (error error) {
	err := store.Save("updo.json", &conf)
	return err
}

func LoadBuiltinApps(conf *Conf) (error error) {
	// TODO storage::LoadBuiltinApps
	return nil
}