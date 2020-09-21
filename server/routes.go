package server

import (
	"github.com/ersonp/go-simple-blockchain/handlers"
	"github.com/ersonp/go-simple-blockchain/middlewares"
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
	Router.GET("/", handlers.ShowIndexPage)

	Router.GET("/home", handlers.ShowHomePage)

	Router.GET("/addblock", handlers.ShowAddBlockPage)

	Router.GET("/nodeinfo", handlers.ShowNodeInfoPage)

	Router.GET("/login", handlers.ShowLoginPage)

	// Group user related routes together
	userRoutes := Router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", middlewares.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", middlewares.EnsureNotLoggedIn(), handlers.PerformLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", middlewares.EnsureLoggedIn(), handlers.Logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", middlewares.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", middlewares.EnsureNotLoggedIn(), handlers.Register)
	}

	// Group article related routes together
	articleRoutes := Router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", handlers.GetArticle)

		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", middlewares.EnsureLoggedIn(), handlers.ShowArticleCreationPage)

		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", middlewares.EnsureLoggedIn(), handlers.CreateArticle)
	}
}
