package main

func main() {
	ConnectDB()
	CreateTable()
	TestModel()
	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "hello"})
	// })
	// router.Run(":8080")
}
