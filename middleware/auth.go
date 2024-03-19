package middleware

import (
	"EasyOps/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, _ := ctx.Get("Token")
		if tokenStr == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 422,
				"msg":  "权限不足",
			})
			ctx.Abort()
		}
		token, myClaim, parseErr := common.ParseToken(ctx, fmt.Sprintf("%v", tokenStr))
		if parseErr != nil || !token.Valid {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 422,
				"msg":  parseErr.Error(),
			})
			ctx.Abort()
		}
		ctx.Set("claim", myClaim)
	}
}
