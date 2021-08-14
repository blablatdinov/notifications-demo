package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Username         []string `json:"username"`
	NotificationText string   `json:"message"`
}

type GetNotificationsInput struct {
	Username string `json:"username"`
}

func (h *Handler) sendNotifications(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{"Некорректный json"})
		return
	}
	id, err := h.services.Notifications.Create(input.Username[0], input.NotificationText)
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetNotifications(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{err.Error()})
	}
	notifications, err := h.services.GetNotifications(userId)
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": notifications,
	})
}

func (h *Handler) DeleteNotification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{"Invalid id param"})
		return
	}
	if err = h.services.Notifications.DeleteNotification("", id); err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{err.Error()})
		return
	}
	c.Status(204)
}

func (h *Handler) GetNotificationsWithUsers(c *gin.Context) {
	notifications, err := h.services.Notifications.GetNotificationsWithUsers()
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": notifications,
	})
}
