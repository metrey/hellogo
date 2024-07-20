package server

import (
    "log"
    "github.com/metrey/hellogo/internal/config"
    "github.com/metrey/hellogo/internal/database"
    "github.com/metrey/hellogo/internal/router"
    "github.com/gin-gonic/gin"
)

func Run() {
    config.LoadConfig()
    db := database.InitDB()
    r := gin.Default()
    router.SetupRoutes(r, db)
    if err := r.Run(); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
