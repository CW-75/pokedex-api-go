package main

import (
	"fmt"
	defaults "pokedex-api-go/src/defaults"
	. "pokedex-api-go/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	envVars := defaults.GetEnvVars()
	SetupRouter(RouterParams{
		Port: envVars.Port,
	})

}
