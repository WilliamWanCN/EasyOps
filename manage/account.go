package manage

import (
	"EasyOps/common"
	"EasyOps/database"
	"EasyOps/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	db := database.GetDb()
	var user = model.User{}
	db.Where("username=?", username).Find(&user)
	if common.DecryptPwd(user.Password) == password {
		tokenStr, err := common.GenToken(ctx, username)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 422,
				"msg":  err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": tokenStr,
		})
		ctx.Set("Token", tokenStr)
		ctx.Next()
		return

	}

}
func Logout(ctx *gin.Context) {

}
func SelfInfo(ctx *gin.Context) {
	db := database.GetDb()
	var user = model.User{}
	username, _ := ctx.Get("username")
	result := db.Where("username=?", username).Find(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  result,
	})
}
