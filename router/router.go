package router

import "github.com/gin-gonic/gin"

func RunServerAndRoutes() {
	r := gin.Default()

	initRoutes(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
