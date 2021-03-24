// article.go

package handlers

import (
	bc "github.com/ersonp/go-simple-blockchain/pkg/blockchain"
	"github.com/ersonp/go-simple-blockchain/pkg/common"
	"github.com/gin-gonic/gin"
)

// ShowBlockchainPage used to show the index page
func ShowBlockchainPage(c *gin.Context) {

	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"title":   "Blockchain Page",
		"payload": bc.Blockchain}, "index.html")
}
