package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	// 配置 Mysql 连接参数
	username := "root"  // 账号
	password := "0801"  // 密码
	host := "127.0.0.1" // ip
	port := 3306        // 端口
	Dbname := "douyin"  // 数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("db connected error", err)
	}
	// 初始化数据库时，创建对应的 token-user 映射表
	NewUserDaoInstance().TokenMap()
	return err
}
