package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"Error": "Handler is not implemented"})
}

func LoginHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"Error": "Handler is not implemented"})
}

func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"Error": "Handler is not implemented"})
}
