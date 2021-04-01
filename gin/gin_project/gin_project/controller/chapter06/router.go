package chapter06

import "github.com/gin-gonic/gin"

func Router(chap06 *gin.RouterGroup)  {

	chap06.GET("/gorm_test",GormTest)

}
