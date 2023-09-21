package router

import "github.com/gin-gonic/gin"

//Initialize is ...
func Initialize() {
	//Initialize router using the Gin default configurations
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)
	//Run the server
	router.Run(":3000")
}
