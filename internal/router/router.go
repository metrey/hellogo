package router

import (
    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    "github.com/metrey/hellogo/internal/handlers"
)

func SetupRoutes(r *gin.Engine, db *sqlx.DB) {
    r.GET("/users", handlers.GetUsers(db))
}
