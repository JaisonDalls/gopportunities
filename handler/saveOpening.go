package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SaveOpeningHandler is ...
func SaveOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Salvo com sucesso!",
	})
}
