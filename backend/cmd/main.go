package main

import (
	"github.com/FashOJ/FashOJ/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
