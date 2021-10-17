package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const path = "~/.config/miladb"

func Load() error {
	if err := checkFolder(); err != nil {
		return err
	}
	if err := loadConfig(); err != nil {
		return err
	}
	return nil
}

func checkFolder() error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fmt.Println(dirname)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("create directory")
		err := os.Mkdir(path, 0655)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("folder exist ")
	}
	return nil
}
func loadConfig() error {
	p := path + "/config.yaml"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(p)
		defer f.Close()
		if err != nil {
			return err
		}
		config := GetDefaultConfig()
		d, err := yaml.Marshal(&config)
		f.Write(d)
	}

	return nil
}
