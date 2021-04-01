package chapter03

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Id int
	Title string
	Desc string
}


func TestSyntaxTpl(ctx *gin.Context)  {

	name := "hallen"

	arr := []int{11,22,33,44}

	article := Article{Id:1,Title:"西游记",Desc:"西游记外传，不容错过"}
	map_data := map[string]interface{}{
		"name":name,
		"arr":arr,
		"birth":"2000-01-02",
		"article":article,
	}
	ctx.HTML(http.StatusOK,"chapter03/test.html",map_data)
}
