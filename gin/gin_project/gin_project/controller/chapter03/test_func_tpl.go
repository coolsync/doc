package chapter03

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
	"html/template"
)

func FuncTpl(ctx *gin.Context)  {


	//fmt.Println("hallen")
	//
	//fmt.Print("hallen")
	//
	//fmt.Printf("name:%s,age:%d\n","hallen",19)
	//
	//
	//name := fmt.Sprint("hallen")
	//fmt.Println(name)
	//
	//user := fmt.Sprintf("name:%s,age:%d,addr:%s\n","hallen",18,"xxx")
	//fmt.Println(user)

	name := "hallen"
	age := 18

	slice_data := []string{"张三","李四","王五"}

	map_test := map[string]interface{}{
		"name":"hallen",
		"age":18,
	}

	now_time := time.Now()
	fmt.Println(now_time)

	map_data := map[string]interface{}{
		"name":name,
		"age":age,
		"slice_data":slice_data,
		"map_test":map_test,
		"now_time":now_time,
	}

	ctx.HTML(http.StatusOK,"chapter03/terst_func_tpl.html",map_data)

}

// 定义函数
func AddNum(num1 int,num2 int) int {

	return num1+num2

}

// qqqqqqqqq     -->  qqqq...

func SubStr(str string,num int) string {

	str_len := len(str)

	// 如果字符串长度小于等于指定的长度，直接返回字符串
	if str_len <= num {
		return str
	}else {  // 要进行加工，然后返回
		new_str := str[0:num]

		return new_str+"..."
	}
}

func Safe(str string) template.HTML {

	return template.HTML(str)
}