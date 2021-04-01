package chapter01

import "github.com/gin-gonic/gin"

func Router(chap01 *gin.RouterGroup)  {

	chap01.GET("/",Hello)

}
