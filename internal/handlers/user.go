package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/metrey/hellogo/internal/repositories"
    "github.com/jmoiron/sqlx"
)

func GetUsers(db *sqlx.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        users, err := repositories.GetAllUsers(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, users)
    }
}
