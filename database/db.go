package database

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb(ctx *gin.Context) *gorm.DB {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("../config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 422,
			"msg":  err.Error(),
		})
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True", viper.Get("db.username"), viper.Get("db.password"), viper.Get("db.host"), viper.Get("db.port"), viper.Get("db.name"))
	if viper.Get("db.type") == "mysql" {
		Db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 422,
				"msg":  err.Error(),
			})
		}
		Db.AutoMigrate()
		return Db
	}
	return Db
}

func GetDb() *gorm.DB {
	return Db
}
