package config

import (
	config "github.com/micro/go-config"
)

// Permission 权限
type Permission struct {
	Service     string `json:"service"`
	Method      string `json:"method"`
	Auth        bool   `json:"auth"`
	Policy      bool   `json:"policy"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Config 配置结构
type Config struct {
	Permissions []Permission `json:"permissions"`
}

// Conf 配置返回
var Conf Config

func init() {
	config.LoadFile("config.yaml")
	config.Scan(&Conf)
}
