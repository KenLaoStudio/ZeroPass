package main

import (
	_ "ZeroPassBackend/docs"
	zeroPassRouter "ZeroPassBackend/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

// @title User API documentation
// @version 1.0.0
// @host localhost:5000
// @BasePath /user
func main() {
	router := gin.Default()

	// Set up Swagger middleware
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Your other routes and handlers
	router.POST("/upload/:address", zeroPassRouter.UploadHandler)
	router.GET("/members", zeroPassRouter.GetAllMembersHandler)
	router.GET("/members/:address", zeroPassRouter.GetMember)
	router.PUT("/members/:address", zeroPassRouter.UpdateMember)
	router.POST("/verify", zeroPassRouter.VerifyHandler)

	_ = router.Run(":8080")
}
