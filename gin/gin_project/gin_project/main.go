package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"gin_project/controller/chapter03"
	"html/template"
	"gin_project/controller/chapter04"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"gin_project/router"
	"gin_project/controller/chapter05"
	_ "gin_project/data_source"
	_ "gin_project/logs_source"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"fmt"
)


func main() {


	router := gin.Default()


	// 加密的盐

	// 基于cookie的存储引擎
	//store := cookie.NewStore([]byte("hallen"))

	// 基于redis的存储引擎
	store,err := redis.NewStore(10,"tcp","192.168.196.129:6379","",[]byte("hallen"))
	fmt.Println(err)
	// 使用session中间件
	router.Use(sessions.Sessions("gin_session",store))

	//router := gin.New()
	//router.Use(gin.Logger(),gin.Recovery())

	// 全局中使用中间件
	//router.Use(chapter05.MiddleWare1)
	router.Use(chapter05.MiddleWare2())
	//router.Use(chapter05.MiddleWare3)


	router.SetFuncMap(template.FuncMap{
		"add":chapter03.AddNum,
		"sub_str":chapter03.SubStr,
		"safe":chapter03.Safe,
	})

	v ,ok := binding.Validator.Engine().(*validator.Validate)

	if ok{
		v.RegisterValidation("len_valid",chapter04.LenValid)
	}


	router.LoadHTMLGlob("template/**/*")
	//router.LoadHTMLFiles("index.html","news.html")

	router.Static("/static","static")
	//router.StaticFS("/static",http.Dir("static"))

	// 路由分组
	all_router.Router(router)

	s := &http.Server{
		Addr:":8080",
		Handler:router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout:5 * time.Second,
	}

	s.ListenAndServe()

}
