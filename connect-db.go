package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	// 设置变量
	user := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "gorm"

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	// 建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: true, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			// 设置表名，列名相关配置
			TablePrefix:   "demo_", // 表名前缀
			SingularTable: false,   // 单数表名
			NoLowerCase:   false,   // 关闭小写转化
		},
		// Logger: SetLoggerConfig(),
	})
	if err != nil {
		log.Fatal(err)
		panic("数据库连接失败, error = " + err.Error())
	}
	DB = db
	fmt.Printf("DB: %v\n", DB)
}
