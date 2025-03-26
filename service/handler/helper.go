package handler

import "github.com/gin-gonic/gin"

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{
		"message": err.Error(),
	})
}

func InternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(500, gin.H{
		"message": err.Error(),
	})
}

func SuccessResponseWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(200, gin.H{
		"message": message,
	})
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}
