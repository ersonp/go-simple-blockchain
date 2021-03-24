package main

import (
	"log"
	"strconv"

	"github.com/ersonp/go-simple-blockchain/server"
	"github.com/gin-gonic/gin"
)

// run is used to start the go gin http server
func run(httpPort int) {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	server.Router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	server.Router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	server.InitializeRoutes()

	// Start serving the application
	log.Println("HTTP Server Listening on port :", httpPort)
	log.Fatal(server.Router.Run(":" + strconv.Itoa(httpPort)))
}
