// article.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/ersonp/go-simple-blockchain/common"
	"github.com/ersonp/go-simple-blockchain/models"
	"github.com/gin-gonic/gin"
)

// ShowIndexPage used to show the index page
func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"title":   "Blockchain Page",
		"payload": articles}, "index.html")
}

// ShowAddBlockPage used to show the index page
func ShowAddBlockPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"title":   "Add Block Page",
		"payload": articles}, "addblock.html")
}

// ShowNodeInfoPage used to show the index page
func ShowNodeInfoPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"title":   "Node Info Page",
		"payload": articles}, "nodeinfo.html")
}

// ShowArticleCreationPage Show article creation page
func ShowArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

// GetArticle is use to get a specific article
func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			common.Render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// CreateArticle to create an article
func CreateArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := models.CreateNewArticle(title, content); err == nil {
		// If the article is created successfully, show success message
		common.Render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
