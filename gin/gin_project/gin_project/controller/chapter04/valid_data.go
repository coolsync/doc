package chapter04

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/astaxie/beego/validation"
)

// binding
//type Article struct {
//	Id int `form:"-"`
//	Title string `form:"title" binding:"len_valid"`
//	Content string `form:"content" binding:"min=5,len_valid"`
//	Desc string `form:"desc"`
//	Uid []User 	`binding:"required,dive,min=10"`
//	//Name [][]int `binding:"min=10,dive,max=20,dive,required"`
//	Name [][]int `binding:"dive,max=20,dive,required"`
//
//}

// valid for beego
type Article struct {
	Id int `form:"-"`
	Title string `form:"title" valid:"Required"`
	Content string `form:"content"`
	Desc string `form:"desc" valid:"Length(6)"`
	Email string `form:"email" valid:"Email"`

}





func ToValidData(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter04/valid_data.html",nil)

}

// binding
//func DoValidData(ctx *gin.Context)  {
//
//	var article Article
//
//	err := ctx.ShouldBind(&article)
//	fmt.Println(err)
//	fmt.Println(article)
//
//	ctx.String(http.StatusOK,"成功")
//
//
//
//}


// beego
func DoValidData(ctx *gin.Context)  {

	var article Article

	err := ctx.ShouldBind(&article)
	fmt.Println(err)
	fmt.Println(article)

	// 初始化验证器



	valid := validation.Validation{}

	var MessageTmpls = map[string]string{
		"Required":     "不能为空",
		"Min":          "不能少于%d个字符",
		"Max":          "不能大于%d个字符",
		"Range":        "取值范围 %d 到 %d",
		"Length":        "长度为:%d",
		"Email":        "邮箱格式不正确",
	}

	validation.SetDefaultMessage(MessageTmpls)


	b,err1 := valid.Valid(&article)

	fmt.Println(err1)

	fmt.Println(len(article.Content))

	key_mapping := map[string]interface{}{
		"Title.Required":"标题",
		"Content.Min":"内容",
		"Desc.Length":"描述",
		"Email.Email":"邮箱",
	}


	if !b {  // 校验有错误
		for _,err2 := range valid.Errors{
			fmt.Println(err2.Key)
			fmt.Println(err2.Message)
			key := key_mapping[err2.Key]
			ctx.String(http.StatusOK, key.(string) + err2.Message)
		}

	}else {
		ctx.String(http.StatusOK, "成功")
	}





}


// 判断是不是大于6
var LenValid validator.Func = func(fl validator.FieldLevel) bool {

	data := fl.Field().Interface().(string)
	fmt.Println("++++++++++++++++++")
	fmt.Println(data)

	fmt.Println(len(data))
	if len(data) > 6{
		return true
	}

	return false
}

