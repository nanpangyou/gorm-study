package main

func main() {
	ConnectDB()
	type User struct {
		ID   int
		Name string
	}
	DB.AutoMigrate(&User{})
	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "hello"})
	// })
	// router.Run(":8080")
}
