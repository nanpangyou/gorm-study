package main

import "fmt"

func TestInsertData() {
	type Insert struct {
		Id    uint    `gorm:"size:10"`
		Name  string  `gorm:"size:25"`
		Age   int     `gorm:"size:3"`
		Email *string `gorm:"size:256"` // 指针类型可以传空值
	}
	// 创建Insert表
	DB.AutoMigrate(&Insert{}) // 创建表

	// 插入单条数据
	email := "1@qq.com"
	err := DB.Create(&Insert{Name: "Jinzhu", Age: 18, Email: &email}).Error
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 批量插入数据
	// 声明一个切片
	var inserts []Insert
	// 循环添加数据
	for i := 0; i < 10; i++ {
		age := 18 + i
		name := fmt.Sprintf("Jinzhu%d", i)
		inserts = append(inserts, Insert{Name: name, Age: age, Email: &email})
	}
	// 批量插入
	err1 := DB.Create(&inserts).Error
	if err1 != nil {
		fmt.Printf("err: %v\n", err1)
	}
}
