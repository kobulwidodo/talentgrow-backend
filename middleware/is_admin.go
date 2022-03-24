package middleware

import (
	"net/http"
	"talentgrow-backend/utils"

	"github.com/gin-gonic/gin"
)

func MustAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := c.MustGet("auth").(CustomClaim)
		if !isAdmin.IsAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewFailResponse("not have access"))
		}
	}
}
