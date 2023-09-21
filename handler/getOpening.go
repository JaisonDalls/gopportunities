package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOpeningHandler is ...
func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Retorno com Sucesso!",
	})
}
