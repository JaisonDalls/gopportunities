package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o Router utilizando as configurações default do gin
	router := gin.Default()
	//Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//rodando a API
	router.Run(":8080")
}
