package main

import (
	"fmt"

	"github.com/pedropassos06/thumbnail-previewer-backend/internal/config"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("Starting server...")
	router.StartServer()
}
