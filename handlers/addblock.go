package handlers

import (
	"net/http"
	"strconv"

	"github.com/ersonp/go-simple-blockchain/blockchain"
	"github.com/ersonp/go-simple-blockchain/common"
	"github.com/gin-gonic/gin"
)

// ShowAddBlockPage used to show the index page
func ShowAddBlockPage(c *gin.Context) {

	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"title":   "Add Block Page",
		"payload": "articles"}, "addblock.html")
}

// HandleWriteBlock used to show the index page
func HandleWriteBlock(c *gin.Context) {
	if bpm, err := strconv.Atoi(c.PostForm("bpm")); err == nil {
		blockchain.AddBlock(bpm)
		c.Redirect(http.StatusFound, "/")
	} else {
		common.Render(c, gin.H{
			"title":        "Add Block Page",
			"ErrorMessage": "please enter a valid numeber"}, "addblock.html")
	}

}
