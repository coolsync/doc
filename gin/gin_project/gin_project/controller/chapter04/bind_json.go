package chapter04

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func ToBindJson(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter04/bind_json.html",nil)

}

func DoBindJson(ctx *gin.Context)  {

	var user User

	err := ctx.ShouldBind(&user)

	fmt.Println(user)
	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code":http.StatusNotFound,
			"msg":"绑定失败",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"msg":"成功",
		})
	}



}
