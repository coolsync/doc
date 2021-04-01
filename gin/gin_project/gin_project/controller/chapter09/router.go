package chapter09

import "github.com/gin-gonic/gin"

func Router(chap09 *gin.RouterGroup)  {

	chap09.GET("/session_test",SessionTest)

}
