package core

import "github.com/SobolevWladimir/miladb/config"

func Start() error {
	err := config.Load()

	return err
}
