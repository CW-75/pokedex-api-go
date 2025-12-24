package routes

import "github.com/gin-gonic/gin"

func routeCallback(c *gin.Context) {
	c.IndentedJSON(200, "Welcome to the API")
}

func SetDefaultRoute(r *gin.Engine) {
	r.GET("/", routeCallback)
}
