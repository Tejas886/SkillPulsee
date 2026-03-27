package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Mock in-memory storage (temporary)
var skills = []map[string]interface{}{}
var idCounter = 1

// GET /api/skills
func GetSkills(c *gin.Context) {
	c.JSON(http.StatusOK, skills)
}

// POST /api/skills
func CreateSkill(c *gin.Context) {
	var skill map[string]interface{}

	if err := c.BindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill["id"] = idCounter
	skill["total_hours"] = 0

	idCounter++
	skills = append(skills, skill)

	c.JSON(http.StatusCreated, skill)
}

// GET /api/skills/:id
func GetSkill(c *gin.Context) {
	id := c.Param("id")

	for _, s := range skills {
		if strconv.Itoa(int(s["id"].(int))) == id {
			c.JSON(http.StatusOK, s)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
}

// DELETE /api/skills/:id
func DeleteSkill(c *gin.Context) {
	id := c.Param("id")

	for i, s := range skills {
		if strconv.Itoa(int(s["id"].(int))) == id {
			skills = append(skills[:i], skills[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
	:}

