package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddResponseToContext(ctx *gin.Context, code int, detail string, data interface{}) {
	ctx.JSON(
		code,
		gin.H{
			"data":    data,
			"status":  code,
			"message": http.StatusText(code),
			"detail":  detail,
		},
	)
	return
}
