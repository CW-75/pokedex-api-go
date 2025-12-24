package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(port int) *gin.Engine {
	r := gin.Default()
	SetDefaultRoute(r)
	return r
}
