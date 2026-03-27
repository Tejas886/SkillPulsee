package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func GetDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total_skills": 0,
		"total_hours":  0,
		"message":      "Dashboard working",
	})
}
