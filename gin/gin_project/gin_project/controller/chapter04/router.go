package chapter04

import "github.com/gin-gonic/gin"

func Router(chap04 *gin.RouterGroup)  {
	// 数据绑定
	chap04.GET("/to_bind_form",ToBindForm)
	chap04.POST("/do_bind_form",DoBingForm)

	chap04.GET("/bind_query_string",BindQueryString)

	chap04.GET("/to_bind_json",ToBindJson)

	chap04.POST("/do_bind_json",DoBindJson)

	chap04.GET("/bind_uri/:name/:age/:addr",BindUri)

	// 数据校验
	chap04.GET("/to_valid",ToValidData)
	chap04.POST("/do_valid",DoValidData)
}
