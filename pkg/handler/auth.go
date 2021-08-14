package handler

import (
	"net/http"

	"github.com/blablatdinov/notifications-demo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input notifications.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{"Некорректный json"})
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input notifications.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{"Некорректный json"})
		return
	}
	token, err := h.services.GenerateToken(input)
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.services.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": users,
	})
}
