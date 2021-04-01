package chapter04

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type User struct {
	Name string `form:"name" json:"name" uri:"name"`
	Age int `form:"age" json:"age" uri:"age"`
	Addr string `form:"addr" json:"addr" uri:"addr"`
}

func ToBindForm(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter04/bing_form.html",nil)

}

func DoBingForm(ctx *gin.Context)  {

	var user User
	err := ctx.ShouldBind(&user)
	fmt.Println(err)

	fmt.Println(user)

	if err != nil {
		ctx.String(http.StatusNotFound,"绑定失败")
	}else {
		ctx.String(http.StatusOK,"绑定成功")
	}





}
