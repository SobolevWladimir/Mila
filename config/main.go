package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const folder = "/.config/miladb"

func Load() error {
	path, err := getPath()
	if err != nil {
		return err
	}
	if err := checkFolder(path); err != nil {
		return err
	}
	if err := loadConfig(path); err != nil {
		return err
	}
	return nil
}

func getPath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dirname + folder, nil
}

func checkFolder(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
func loadConfig(path string) error {
	p := path + "/config.yaml"

	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		fmt.Println("create default config: " + p)
		f, err := os.Create(p)
		defer f.Close()
		if err != nil {
			return err
		}
		config := GetDefaultConfig()
		d, err := yaml.Marshal(&config)
		f.Write(d)
	} else {
		fmt.Println("load config: " + p)
	}

	return nil
}
