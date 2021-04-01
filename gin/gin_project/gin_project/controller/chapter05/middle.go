package chapter05

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)

func MiddleWare1(ctx *gin.Context)  {
	start_time := time.Now()

	fmt.Println("这是自定义中间件1--开始")
	ctx.Next()

	time_count := time.Since(start_time)
	fmt.Println(time_count)
	fmt.Println("这是自定义中间件1--结束")
}

func MiddleWare2() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		fmt.Println("这是自定义中间件2--开始")

		//if 3 < 4{   // 满足条件
		//	ctx.Abort()
		//}
		time.Sleep(3)
		ctx.Next()
		fmt.Println("这是自定义中间件2--结束")
	}
}

func MiddleWare3(ctx *gin.Context)  {
	fmt.Println("这是自定义中间件3--开始")
	time.Sleep(3)
	ctx.Next()
	fmt.Println("这是自定义中间件3--结束")
}


// 1-开始 --》 2--开始 -- 》3-开始 -- 3-结束--》2--结束--》1-结束

