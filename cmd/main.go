package main

import "github.com/gin-gonic/gin"
import zeroPassRouter "ZeroPassBackend/router"

func main() {
	router := gin.Default()

	router.POST("/upload", zeroPassRouter.UploadHandler)
	router.POST("/verify", zeroPassRouter.VerifyHandler)

	_ = router.Run(":8080")
}
