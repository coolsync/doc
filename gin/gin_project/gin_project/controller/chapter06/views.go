package chapter06

import (
	"github.com/gin-gonic/gin"

	"gin_project/models"
	"gin_project/data_source"
)

func GormTest(ctx *gin.Context)  {

	user := models.User{Name:"hallen",Age:18}

	data_source.Db.Create(&user)

	data_source.Db.Close()
}
