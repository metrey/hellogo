package utils

import (
    "log"
    "github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, message string) {
    log.Printf("%s: %v", message, err)
    c.JSON(500, gin.H{"error": message})
}
