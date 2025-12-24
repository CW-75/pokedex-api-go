package defaults

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EnvVars struct {
	GinMode string
	Port    int
}

func getEnvPort() int {
	portStr := os.Getenv("PORT")
	port, error := strconv.Atoi(portStr)
	if error != nil {
		fmt.Println("Port was not found, will be set Port=8080")
		return 8080 // Default port if not set
	}
	return port
}

func GetEnvVars() EnvVars {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug" // Default mode
	}
	gin.SetMode(ginMode)
	port := getEnvPort()
	if ginMode == "release" {
		fmt.Println("Running in", ginMode, "mode", "on port:", port)
	}
	return EnvVars{
		GinMode: ginMode,
		Port:    port,
	}
}
