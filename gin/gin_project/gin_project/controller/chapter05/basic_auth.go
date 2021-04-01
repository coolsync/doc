package chapter05

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// 私有数据

var map_data map[string]interface{} = map[string]interface{}{
	"zs": gin.H{"age":18,"addr":"zs-xx"},
	"ls": gin.H{"age":19,"addr":"ls-xx"},
	"ww": gin.H{"age":20,"addr":"ww-xx"},
}


func AuthTest(ctx *gin.Context)  {
	user_name := ctx.Query("user_name")
	data,ok := map_data[user_name]

	if ok {
		ctx.JSON(http.StatusOK,gin.H{"user":user_name,"data":data})
	}else {
		ctx.JSON(http.StatusNotFound,gin.H{"user":user_name,"data":"没有数据"})
	}
}

func WrapFTest(w http.ResponseWriter,r *http.Request)  {
	
}
