package utils

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, status int, payload any) {
	c.JSON(status, payload)
}
