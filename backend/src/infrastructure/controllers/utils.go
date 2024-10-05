package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPathParams[T any](ctx *gin.Context) (pathParams T, ok bool) {
	err := ctx.ShouldBindUri(&pathParams)
	ok = err == nil
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
	return
}