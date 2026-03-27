package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSkills(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetSkills working",
	})
}

func CreateSkill(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "CreateSkill working",
	})
}

func GetSkill(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetSkill working",
	})
}

func DeleteSkill(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteSkill working",
	})
}
