package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Host string
	Port int
	Name string
}

func Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "rpcui",
	})
}
