package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type UserInfo struct {
	Id int `form:"id"`
	Name string `form:"name"`
	Age int `form:"age"`
	Addr string `form:"addr"`

}


func User(ctx *gin.Context)  {

	name := "hallen"
	ctx.HTML(200,"user/user.html",name)
}


// 结构体渲染
func UserInforStruct(ctx *gin.Context)  {

	//user_info := UserInfo{Id:1,Name:"hallen",Age:19,Addr:"xxxx"}

	var user_info1 UserInfo
	user_info1.Id = 2
	user_info1.Name = "hallen2"
	user_info1.Age = 18
	user_info1.Addr = "xxx"



	ctx.HTML(http.StatusOK,"chapter02/user_info.html",user_info1)

}

// 数组
func ArrController(ctx *gin.Context)  {

	arr := [3]int{1,3,5}

	ctx.HTML(http.StatusOK,"chapter02/arr.html",arr)

}

// 结构体数组
func ArrStruct(ctx *gin.Context)  {

	arr_statuct := [3]UserInfo{{Id:1,Name:"hallen",Age:18,Addr:"xxx"},{Id:2,Name:"hallen2",Age:19,Addr:"xxx2"},{Id:3,Name:"hallen3",Age:20,Addr:"xxx3"}}
	ctx.HTML(http.StatusOK,"chapter02/arr_struct.html",arr_statuct)

}

// map
func MapController(ctx *gin.Context)  {

	map_data := map[string]string{"name":"hallen","age":"18"}

	map_data2 := map[string]int{"age":18}

	map_data3 := map[string]interface{}{"map_data":map_data,"map_data2":map_data2}

	ctx.HTML(http.StatusOK,"chapter02/map.html",map_data3)

}

// map + struct
func MapStruct(ctx *gin.Context)  {

	map_struct := map[string]UserInfo{"user":{Id:1,Name:"hallen",Age:18,Addr:"xxx"}}
	ctx.HTML(http.StatusOK,"chapter02/map_struct.html",map_struct)

}

// 切片

func SliceController(ctx *gin.Context)  {

	slice_data := []int{1,2,3,4,5,6,7,9}
	ctx.HTML(http.StatusOK,"chapter02/slice.html",slice_data)
}


// 路径中直接加上参数值

func Param1(ctx *gin.Context)  {

	id := ctx.Param("id")

	ctx.String(http.StatusOK,"hello,%s",id)
}

func Param2(ctx *gin.Context)  {

	id := ctx.Param("id")

	ctx.String(http.StatusOK,"hello,%s",id)
}

func GetQueryData(ctx *gin.Context)  {
	id := ctx.Query("id")

	name := ctx.DefaultQuery("name","hallen")
	fmt.Println(name)
	ctx.String(http.StatusOK,"hallen,%s",id)

}


func GetQueryArrData(ctx *gin.Context)  {
	ids := ctx.QueryArray("ids")
	ctx.String(http.StatusOK,"hallen,%s",ids)

}

func GetQueryMapData(ctx *gin.Context)  {

	user := ctx.QueryMap("user")
	ctx.String(http.StatusOK,"hallen,%s",user)

}


// 到用户添加页面
func ToUserAdd(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter02/user_add.html",nil)
}

// 添加用户
func DoUserAdd(ctx *gin.Context)  {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)

	ctx.String(http.StatusOK,"添加成功")

}


// 到用户添加页面
func ToUserAdd2(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter02/user_add2.html",nil)
}


// 添加用户
func DoUserAdd2(ctx *gin.Context)  {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.DefaultPostForm("age","18")

	user := ctx.PostFormMap("user")

	loves := ctx.PostFormArray("love")
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(age)
	fmt.Println(loves)
	fmt.Println(user)

	ctx.String(http.StatusOK,"添加成功")

}

// 到用户添加页面
func ToUserAdd3(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter02/user_add3.html",nil)
}

func DoUserAdd3(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.PostForm("age")

	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(age)
	//map_data := map[string]interface{}{
	//	"code":200,
	//	"msg":"成功",
	//}
	ctx.JSON(http.StatusOK,gin.H{"code":200,"msg":"成功"})

}


// 到用户添加页面
func ToUserAdd4(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter02/user_add4.html",nil)
}

func DoUserAdd4(ctx *gin.Context)  {

	var user_info UserInfo

	err := ctx.ShouldBind(&user_info)
	fmt.Println(err)
	fmt.Println(user_info)

	ctx.String(http.StatusOK,"绑定成功")


}

