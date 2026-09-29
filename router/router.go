package router

import "github.com/gin-gonic/gin"

func Inicialiaze() {
	r := gin.Default()
	InicializeRoutes(r)
	r.Run(":8080")
}
