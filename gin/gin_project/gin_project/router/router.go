package all_router

import (
	"github.com/gin-gonic/gin"
	"gin_project/controller/chapter01"
	"gin_project/controller/chapter02"
	"gin_project/controller/chapter03"
	"gin_project/controller/chapter04"
	"gin_project/controller/chapter05"
	"gin_project/controller/chapter06"
	"gin_project/controller/chapter07"
	"gin_project/controller/chapter09"
)

func Router(router *gin.Engine)  {

	chap01 := router.Group("/chapter01")
	chap02 := router.Group("/chapter02")
	chap02.Use(chapter05.MiddleWare1)
	chap03 := router.Group("/chapter03")
	chap04 := router.Group("/chapter04")
	chap05 := router.Group("/chapter05")
	chap06 := router.Group("/chapter06")
	chap07 := router.Group("/chapter07")
	chap09 := router.Group("/chapter09")

	chapter01.Router(chap01)
	chapter02.Router(chap02)
	chapter03.Router(chap03)
	chapter04.Router(chap04)
	chapter05.Router(chap05)
	chapter06.Router(chap06)
	chapter07.Router(chap07)
	chapter09.Router(chap09)



}
