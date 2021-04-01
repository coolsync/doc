package chapter03

import "github.com/gin-gonic/gin"

func Router(chap03 *gin.RouterGroup)  {
	// 模板语法
	chap03.GET("/test_tpl",TestSyntaxTpl)
	chap03.GET("/test_func_tpl",FuncTpl)
}
