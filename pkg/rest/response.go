package rest

import (
	"github.com/DiamondDmitriy/big-note-api/pkg/rest/types"
	"github.com/gin-gonic/gin"
)

func newResponse() {
}

func ResponseError(ctx *gin.Context, code int, message string, details interface{}) {
	ctx.JSON(code, types.ResponseError{
		Status:  "error",
		Message: message,
		Details: details,
	})
}

func ResponseSuccess(ctx *gin.Context, code int, data interface{}, meta interface{}) {
	ctx.JSON(code, types.ResponseSuccess{
		Status: "success",
		Data:   data,
		Meta:   meta,
	})
}
