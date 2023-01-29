package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	// 连接数据库：name、password、ip、port、库名
	// 使用mysql驱动
	dsn := "root:admin@tcp(82.157.130.123:33306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
