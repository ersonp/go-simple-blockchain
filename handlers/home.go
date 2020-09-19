package handlers

import (
	"github.com/ersonp/go-simple-blockchain/common"
	"github.com/ersonp/go-simple-blockchain/models"
	"github.com/gin-gonic/gin"
)

func ShowHomePage(c *gin.Context) {
	// Call the render function with the name of the template to render
	education := models.GetEducationList()
	experience := models.GetExperienceList()
	profession := models.GetProfessionList()
	project := models.GetProjectList()
	myinfo := models.GetMyInfoData()
	common.Render(c, gin.H{
		"title":      "Home",
		"MyInfo":     myinfo,
		"Profession": profession,
		"Education":  education,
		"Experience": experience,
		"Project":    project,
	}, "home.html")
}
