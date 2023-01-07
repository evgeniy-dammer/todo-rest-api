package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// myError
type myError struct {
	Message string `json:"message"`
}

// newErrorResponse is a response with error
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, myError{Message: message})
}
