package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/dashboard
func GetDashboard(c *gin.Context) {
	totalSkills := len(skills)
	totalHours := 0.0

	for _, s := range skills {
		totalHours += s["total_hours"].(float64)
	}

	c.JSON(http.StatusOK, gin.H{
		"total_skills": totalSkills,
		"total_hours":  totalHours,
	})
}

// GET /health
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

