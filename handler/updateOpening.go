package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateOpeningHandler is ...
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Atualizado com Sucesso!",
	})
}
