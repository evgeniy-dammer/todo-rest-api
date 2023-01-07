package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// createList
func (h *Handler) createList(c *gin.Context) {
	// just for authorization testing
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// getAllLists
func (h *Handler) getAllLists(c *gin.Context) {

}

// getListById
func (h *Handler) getListById(c *gin.Context) {

}

// updateList
func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
