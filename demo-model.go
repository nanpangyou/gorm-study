package main

func TestModel() {
	// 定义结构体
	type Student struct {
		Id    uint    `gorm:"size:10"`
		Name  string  `gorm:"size:25"`
		Age   int     `gorm:"size:3"`
		Email *string `gorm:"size:256"` // 指针类型可以传空值
	}
	// 可以通过AutoMigrate来生成表
	// AutoMigrate的逻辑是只新增，不删除，不修改（字段大小会修改）
	// 例如将Name字段改为Name1，则会多出一个Name1字段
	DB.AutoMigrate(&Student{})

	// 还有其他可使用的属性, 多属性之间用分号(;)隔开
	type Demostruct struct {
		Id    uint    `gorm:"size:10"`
		Name  string  `gorm:"type:varchar(10);not null;comment:姓名"`
		Age   int     `gorm:"size:3;comment:年龄"`
		Email *string `gorm:"size:256;unique;comment:邮箱"` // 指针类型可以传空值
	}
	DB.AutoMigrate(&Demostruct{})
}
