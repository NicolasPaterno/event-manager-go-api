package middlewares

import (
	"net/http"

	"github.com/nicao/minimal-goapi/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unathorized"})
		return
	}
	userid, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unathorized"})
		return
	}

	context.Set("userId", userid)
	context.Next()
}
