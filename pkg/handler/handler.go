package handler

import (
	"github.com/blablatdinov/notifications-demo/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api/v1", h.userIdentityFromHeader)
	{
		notifications := api.Group("/notifications")
		{
			notifications.GET("/user/", h.GetNotifications)
			notifications.GET("/", h.GetNotificationsWithUsers)
			notifications.POST("/", h.sendNotifications)
			notifications.DELETE("/:id", h.DeleteNotification)
		}
		users := api.Group("/users")
		{
			users.GET("/", h.GetUsers)
		}
	}
	ws := router.Group("/ws", h.userIdentityFromQuery)
	{
		ws.Any("/", h.wsGinHandler)
	}
	return router
}
