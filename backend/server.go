package main

import (
	"os"
	"commercial-shop.com/routes"
)

func main() {
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "4505"
	}
	routes.Run(port)
}
