package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type RouterParams struct {
	Addr string
	Port int
}

func SetupRouter(params RouterParams) (err error) {
	r := gin.Default()
	setDefaultRoute(r)
	return r.Run(params.Addr + ":" + strconv.Itoa(params.Port))
}
