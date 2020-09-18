package main

import (
	"net/http"

	"rizhua.com/rpcui/api"

	"github.com/gin-gonic/gin"
)

// import "rizhua.com/rpcui/router"

func main() {

	router := gin.Default()
	router.StaticFS("/assets", http.Dir("D:/www/go/src/rizhua.com/rpcui/assets"))
	router.LoadHTMLGlob("view/*")
	router.GET("/", api.Index)

	router.POST("/send", api.Send)

	// 指定地址和端口号
	router.Run("127.0.0.1:8080")
}
