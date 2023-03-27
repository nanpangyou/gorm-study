package main

import (
	"fmt"

	"gorm.io/gorm"
)

var selectSession *gorm.DB

func TestSelectData() {
	selectSession = DB.Session(&gorm.Session{Logger: SetLoggerConfig()})
	selectSingleData()
	selectMultiData()
}

func selectSingleData() {
	// 查询单条数据
	var demostruct Demostruct
	selectSession.Take(&demostruct) // 查询第一条数据，相当于执行 SELECT * FROM `demo_demostructs` LIMIT 1
	fmt.Printf("demostruct.take: %v\n", demostruct)
	fmt.Printf("demostruct.Email: %v\n", *demostruct.Email)
	demostruct = Demostruct{}
	selectSession.First(&demostruct) // 查询第一条数据，相当于执行 SELECT * FROM `demo_demostructs` ORDER BY `demo_demostructs`.`id` LIMIT 1
	fmt.Printf("demostruct.first: %v\n", demostruct)
	fmt.Printf("demostruct.Email: %v\n", *demostruct.Email)
	demostruct = Demostruct{}
	selectSession.Last(&demostruct) // 查询最后一条数据，相当于执行 SELECT * FROM `demo_demostructs` ORDER BY `demo_demostructs`.`id` DESC LIMIT 1
	fmt.Printf("demostruct.last: %v\n", demostruct)
	fmt.Printf("demostruct.Email: %v\n", *demostruct.Email)
}

func selectMultiData() {
	// 查询多条数据

}
