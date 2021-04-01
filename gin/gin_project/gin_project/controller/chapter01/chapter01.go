package chapter01

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context)  {
	//ctx.String(200,"hello gin")

	name := "hallen"
	ctx.HTML(200,"chapter01/index.html",name)
}
