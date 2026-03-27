package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/trainwithshubham/skillpulse/database"
    "github.com/trainwithshubham/skillpulse/handlers"
)

func main() {
    database.Connect()

    router := gin.Default()

    api := router.Group("/api")
    {
        api.GET("/skills", handlers.GetSkills)
        api.POST("/skills", handlers.CreateSkill)
        api.GET("/skills/:id", handlers.GetSkill)
        api.DELETE("/skills/:id", handlers.DeleteSkill)
        api.POST("/skills/:id/log", handlers.CreateLog)
        api.GET("/dashboard", handlers.GetDashboard)
    }

    router.GET("/health", handlers.HealthCheck)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("SkillPulse API running on port %s", port)
    router.Run(":" + port)
}
