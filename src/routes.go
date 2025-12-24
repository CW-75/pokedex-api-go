package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(port int) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the PokeAPI")
	})

	return r
}
