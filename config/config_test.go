package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println(
		Conf.Permissions[0].Service,
		Conf.Permissions[0].Method,
		Conf.Permissions[0].Auth,
		Conf.Permissions[0].Policy,
		Conf.Permissions[0].Name,
		Conf.Permissions[0].Description,
	)
}
