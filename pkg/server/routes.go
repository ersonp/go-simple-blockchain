package server

import (
	"github.com/ersonp/go-simple-blockchain/pkg/handlers"
	"github.com/ersonp/go-simple-blockchain/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

// Router is defined here
var Router *gin.Engine

// InitializeRoutes is used to initilies routes
func InitializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	Router.Use(middlewares.SetUserStatus())

	// Set the path for the static folder
	middlewares.SetStatic(Router)

	// Handle the index route
	Router.GET("/", handlers.ShowBlockchainPage)
	Router.GET("/addblock", handlers.ShowAddBlockPage)
	Router.POST("/add", handlers.HandleWriteBlock)
	Router.GET("/nodeinfo", handlers.ShowNodeInfoPage)
}
