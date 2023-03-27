package main

func main() {

	// 链接数据库
	ConnectDB()

	// 创建表
	// CreateTable()

	// 创建表（更多的条件）
	// TestModel()

	// 测试插入单条数据或者批量插入数据
	// TestInsertData()

	// 测试查询数据
	TestSelectData()

	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "hello"})
	// })
	// router.Run(":8080")
}
