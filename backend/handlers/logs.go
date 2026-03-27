package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trainwithshubham/skillpulse/database"
)

// POST /api/skills/:id/log
func CreateLog(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Hours   float64 `json:"hours"`
		Notes   string  `json:"notes"`
		LogDate string  `json:"log_date"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.LogDate == "" {
		body.LogDate = time.Now().Format("2006-01-02")
	}

	_, err := database.DB.Exec(`
		INSERT INTO learning_logs (skill_id, hours, notes, log_date)
		VALUES (?, ?, ?, ?)
	`, id, body.Hours, body.Notes, body.LogDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log added"})
}
