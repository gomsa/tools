package test

import (
	"log"
	"testing"

	db "github.com/gomsa/tools/database"
)

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
func TestDB(t *testing.T) {
	db, err := db.Connection(&db.Config{
		Driver: "mysql",
		// Host 主机地址
		Host: "127.0.0.1",
		// Port 主机端口
		Port: "3306",
		// User 用户名
		User: "root",
		// Password 密码
		Password: "1234567",
		// DbName 数据库名称
		DbName: "user_service",
	})
	if err != nil {
		t.Errorf("Database connection failed, %v!", err)
	}
	log.Println(db, err)
	// if actual != expected {
	// 	t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)

	// }
}
