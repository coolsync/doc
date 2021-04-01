package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin_project/proto"
	"fmt"
)

func OutJson(ctx *gin.Context)  {

	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"tag":"<br>",
		"msg":"提交成功",
		"html":"<b>Hello, world!</b>",
	})

}

func OutAsciiJson(ctx *gin.Context)  {

	ctx.AsciiJSON(http.StatusOK,gin.H{
		"code":200,
		"tag":"<br>",
		"msg":"提交成功",
		"html":"<b>Hello, world!</b>",
	})

}


func OutJsonp(ctx *gin.Context)  {

	ctx.JSONP(http.StatusOK,gin.H{
		"code":200,
		"tag":"<br>",
		"msg":"提交成功",
		"html":"<b>Hello, world!</b>",
	})

}

func OutPureJsonp(ctx *gin.Context)  {

	ctx.PureJSON(http.StatusOK,gin.H{
		"code":200,
		"tag":"<br>",
		"msg":"提交成功",
		"html":"<b>Hello, world!</b>",
	})

}

func OutSecureJsonp(ctx *gin.Context)  {

	ctx.SecureJSON(http.StatusOK,gin.H{
		"code":200,
		"tag":"<br>",
		"msg":"提交成功",
		"html":"<b>Hello, world!</b>",
	})

}

func OutXml(ctx *gin.Context)  {

	ctx.XML(http.StatusOK,gin.H{
		"code":200,
		"tag":"<br>",
		"msg":"提交成功",
		"html":"<b>Hello, world!</b>",
	})

}

func OutYaml(ctx *gin.Context)  {

	ctx.YAML(http.StatusOK,gin.H{
		"code":200,
		"tag":"<br>",
		"msg":"提交成功",
		"html":"<b>Hello, world!</b>",
	})

}

func OutProto(ctx *gin.Context)  {

	user_data := &user.User{Name:"hallen",Age:18}

	ctx.ProtoBuf(http.StatusOK,user_data)

}

func RedirectA(ctx *gin.Context)  {
	fmt.Println("这是A路由")
	ctx.Redirect(http.StatusFound,"/redirect_b")
	//ctx.Redirect(http.StatusFound,"https://study.163.com/course/courseMain.htm?courseId=1210678804&share=2&shareId=1025897964")
}

func RedirectB(ctx *gin.Context)  {
	fmt.Println("这是B路由")

	ctx.String(http.StatusOK,"这是B路由")

}
