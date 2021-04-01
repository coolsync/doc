package chapter04

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func BindUri(ctx *gin.Context)  {

	var user User

	err := ctx.ShouldBindUri(&user)

	fmt.Println(user)

	if err !=nil {
		ctx.String(http.StatusNotFound,"绑定失败")
	}

	ctx.String(http.StatusOK,"绑定成功")



}
