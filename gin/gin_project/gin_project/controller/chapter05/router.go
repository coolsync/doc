package chapter05

import "github.com/gin-gonic/gin"

func Router(chap05 *gin.RouterGroup)  {

	chap05.GET("/auth_test",gin.BasicAuth(gin.Accounts{
		"zs":"123456",
		"ls":"123",
		"www":"1234",
	}),gin.WrapF(WrapFTest),AuthTest)

}
