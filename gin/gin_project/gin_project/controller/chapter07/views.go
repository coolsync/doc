package chapter07

import (
	"github.com/gin-gonic/gin"
	"gin_project/logs_source"
)

func LogTest(ctx *gin.Context)  {

	logs_source.Log.Info("这是info級別")

	logs_source.Log.Debug("这是debug級別")



}
