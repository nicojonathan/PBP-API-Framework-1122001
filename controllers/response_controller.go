package controllers

import (
	m "belajar_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendUserSuccessResponse(c *gin.Context, message string, data []m.User) {
    response := gin.H{
        "status":  http.StatusOK,
        "message": message,
        "data":    data,
    }

    c.JSON(http.StatusOK, response)
}

func sendSuccessResponse(c *gin.Context, message string) {
    response := gin.H{
        "status":  http.StatusOK,
        "message": message,
    }

    c.JSON(http.StatusOK, response)
}


func sendErrorResponse(c *gin.Context, status int, message string) {
    response := gin.H{
        "status":  status,
        "message": message,
    }

    c.JSON(status, response)
}