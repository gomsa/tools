package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println(
		Conf.App,
		Conf.Version,
		Conf.Permissions[0].Service,
		Conf.Permissions[0].Method,
		Conf.Permissions[0].Auth,
		Conf.Permissions[0].Policy,
		Conf.Permissions[0].DisplayName,
		Conf.Permissions[0].Description,
	)
}
