package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"time": time.Now()})
}
