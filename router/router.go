package router

import "github.com/gin-gonic/gin"

func UploadHandler(c *gin.Context) {
	// 解析上傳的文件
	// 驗證文件
	// 將文件存儲到 MongoDB
	// 返回結果
}

func VerifyHandler(c *gin.Context) {
	// 驗證 ZKP credential
	// 與智能合約交互
	// 返回結果
}
