package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	conf := &Config{}
	conf.LoadFile("config.yaml")
	App := conf.GetApp()
	Version := conf.GetVersion()
	Permissions := conf.GetPermissions()
	fmt.Println(
		App,
		Version,
		Permissions[0].Service,
		Permissions[0].Method,
		Permissions[0].Auth,
		Permissions[0].Policy,
		Permissions[0].DisplayName,
		Permissions[0].Description,
	)
}
