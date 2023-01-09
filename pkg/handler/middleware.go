package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

// userIdentity validate access token
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	
	headerParts := strings.Split(header, " ")
	
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	
	// parse token and return user id
	userId, err := h.services.ParseToken(headerParts[1])
	
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	
	c.Set(userCtx, userId)
}

// getUserId returns user id from authorization context
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user is not found")
	}
	
	// interface to int casting
	idInt, ok := id.(int)
	
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user is not found")
	}
	
	return idInt, nil
}
