package xorm

import (
	"testing"

	xorm "github.com/gomsa/tools/xorm"
)

func TestXorm(t *testing.T) {
	_, err := xorm.Connection(&xorm.Config{
		Driver: "odbc",
		// Host 主机地址
		Host: "192.168.1.26",
		// Port 主机端口
		Port: "1433",
		// User 用户名
		User: "sa",
		// Password 密码
		Password: "",
		// DbName 数据库名称
		DbName: "bvbv01",
		// Charset 数据库编码
		Charset: "gbk",
	})
	if err != nil {
		t.Errorf("Database connection failed, %v!", err)
	}
}
