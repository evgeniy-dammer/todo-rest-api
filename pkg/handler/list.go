package handler

import (
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getAllListsResponse
type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// createList is a new list creation handler
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	var input todo.TodoList

	// get data from body
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// getAllLists is a get all lists handler
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}

// getListById is a get list by id handler
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// updateList is an update list by id handler
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.UpdateById(userId, listId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

// deleteList is a delete list by id handler
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.DeleteById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
