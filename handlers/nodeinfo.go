// user.go

package handlers

import (
	"fmt"

	bc "github.com/ersonp/go-simple-blockchain/blockchain"
	"github.com/ersonp/go-simple-blockchain/common"
	"github.com/gin-gonic/gin"
)

// ShowNodeInfoPage used to show the index page
func ShowNodeInfoPage(c *gin.Context) {

	var a bc.Address

	a = bc.HostAddressField

	fmt.Println(bc.PeerMetrics)
	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"NodeId":      a.HostID,
		"NodeAddress": a.FullAddr,
		"PeerData":    bc.PeerMetrics,
		"payload":     "articles"}, "nodeinfo.html")
}
