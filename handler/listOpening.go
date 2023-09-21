package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListOpeningsHandler is ...
func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Retorno da lista com Sucesso!",
	})
}
