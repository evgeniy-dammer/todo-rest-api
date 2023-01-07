package handler

import (
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// signUp
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(
		http.StatusOK,
		map[string]interface{}{"id": id},
	)
}

// signIn
func (h *Handler) signIn(c *gin.Context) {

}
