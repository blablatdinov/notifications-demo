package handler

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context, token string) {
	userId, username, err := h.services.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, ErrorResponse{err.Error()})
		return
	}
	c.Set(userCtx, userId)
	c.Set("username", username)
}

func (h *Handler) userIdentityFromHeader(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(401, ErrorResponse{"empty auth header"})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(401, ErrorResponse{"invalid auth header"})
		return
	}
	h.userIdentity(c, headerParts[1])
}

func (h *Handler) userIdentityFromQuery(c *gin.Context) {
	token, ok := c.GetQuery("token")
	if !ok {
		c.AbortWithStatusJSON(401, ErrorResponse{"invalid auth query param"})
	}
	h.userIdentity(c, token)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getUsername(c *gin.Context) (string, error) {
	username, ok := c.Get("username")
	if !ok {
		return "", errors.New("username not found")
	}

	usernameString, ok := username.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return usernameString, nil
}
