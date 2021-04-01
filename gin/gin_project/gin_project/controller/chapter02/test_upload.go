package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
	"strconv"
)

// form表单单文件上传
func ToUpload1(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter02/test_upload1.html",nil)
}

func DoUpload1(ctx *gin.Context)  {
	file,_ := ctx.FormFile("file")

	fmt.Println(file.Filename)

	time_unix_int  := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int,10)
	dst := "upload/" + time_unix_str + file.Filename
	ctx.SaveUploadedFile(file,dst)


	ctx.String(http.StatusOK,"上传成功")

}

// form表单多文件上传
func ToUpload2(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"chapter02/test_upload2.html",nil)
}



func DoUpload2(ctx *gin.Context)  {

	form,_ := ctx.MultipartForm()
	files := form.File["file"]

	for _,file := range files {
		fmt.Println(file.Filename)
		time_unix_int  := time.Now().Unix()
		time_unix_str := strconv.FormatInt(time_unix_int,10)
		dst := "upload/" + time_unix_str + file.Filename
		ctx.SaveUploadedFile(file,dst)
	}


	ctx.String(http.StatusOK,"上传成功")

}


// ajax单文件

func ToUploadFile3(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"chapter02/test_upload3.html",nil)
}

func DoUploadFile3(ctx *gin.Context)  {

	name := ctx.PostForm("name")
	fmt.Println(name)

	file,_ := ctx.FormFile("file")

	fmt.Println(file.Filename)

	time_unix_int  := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int,10)
	dst := "upload/" + time_unix_str + file.Filename
	ctx.SaveUploadedFile(file,dst)

	ctx.String(http.StatusOK,"上传成功")


}
// ajax多文件

func ToUploadFile4(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"chapter02/test_upload4.html",nil)
}

func DoUploadFile4(ctx *gin.Context)  {
	form,_ := ctx.MultipartForm()
	
	files := form.File["file"]

	name := ctx.PostForm("name")
	fmt.Println(name)
	
	for _, file := range files{
		fmt.Println(file.Filename)
		time_unix_int  := time.Now().Unix()
		time_unix_str := strconv.FormatInt(time_unix_int,10)
		dst := "upload/" + time_unix_str + file.Filename
		ctx.SaveUploadedFile(file,dst)
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"上传成功",
	})
}