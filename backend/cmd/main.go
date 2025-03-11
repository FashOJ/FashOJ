package main

import (
	"FashOJ/backend/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
