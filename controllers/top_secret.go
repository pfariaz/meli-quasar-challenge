package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessMessageLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world from ProcessMessageLocation"})
}

func ProcessPartialMessageLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world from ProcessPartialMessageLocation"})
}

func GetPartialMessageLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world from GetPartialMessageLocation"})
}
