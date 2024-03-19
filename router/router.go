package router

import (
	"EasyOps/manage"
	"EasyOps/middleware"

	"github.com/gin-gonic/gin"
)

func RunServer() *gin.Engine {
	g := gin.Default()
	index := g.Group("/")
	{
		index.POST("/login", manage.Login)
		index.GET("/logout", middleware.AuthMiddleware(), manage.Logout)
	}
	account := g.Group("/account")
	{
		account.GET("/selfInfo", middleware.AuthMiddleware(), manage.SelfInfo)
	}
	sys := g.Group("/sys")
	{
	}
	return g
}
