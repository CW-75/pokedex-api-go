package main

import (
	"fmt"
	. "pokedex-api-go/src/routes"
)

func main() {
	fmt.Println("Hello, World")
	SetupRouter(8080).Run()
}
