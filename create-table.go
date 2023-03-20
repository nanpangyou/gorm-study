package main

type User struct {
	ID   int
	Name string
}

func CreateTable() {

	// gorm默认的命名策略是，表名是蛇形复数，字段是蛇形单数
	// 上面的结构会生成以下结构
	// create table `users` (`id` bigint,`name` longtext)
	// 也可以在创建连接的时候设置参数来自定义
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "demo_", // 表名前缀
	// 		SingularTable: false,   // 单数表名
	// 		NoLowerCase:   false,   // 关闭小写转化
	//  },

	// 创建user表
	DB.AutoMigrate(&User{})
	// 创建数据
	DB.Create(&User{Name: "Jinzhu"})
	TestLoggerConfig()
}
