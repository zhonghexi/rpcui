package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// SuccessCode 成功返回码
	SuccessCode int = 10000
	// SuccessDesc 成功返回描述
	SuccessDesc string = "Success"
	// FailCode 失败返回码
	FailCode int = 20000
	// FailDesc 失败返回描述
	FailDesc string = "Fail"
)

// Ok ...
func Ok(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"desc": SuccessDesc,
	})
}

// OkHasInfo ...
func OkHasInfo(info interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"info": info,
		"desc": SuccessDesc,
	})
}

// OkHasList ...
func OkHasList(list interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"list": list,
		"desc": SuccessDesc,
	})
}

// OkHasPage ...
func OkHasPage(total int64, list gin.H, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":  SuccessCode,
		"list":  list,
		"total": total,
		"desc":  SuccessDesc,
	})
}

// Fail ...
func Fail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": FailCode,
		"desc": FailDesc,
	})
}

// FailHasCode ...
func FailHasCode(code int, desc string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"desc": desc,
	})
}

// FailHasDesc ...
func FailHasDesc(desc string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": FailCode,
		"desc": desc,
	})
}
