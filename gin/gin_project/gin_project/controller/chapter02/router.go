package chapter02

import (
	"github.com/gin-gonic/gin"
	"gin_project/controller/chapter05"
)

func Router(chap02 *gin.RouterGroup)  {

	chap02.GET("/user",User)

	chap02.GET("/user_struct",chapter05.MiddleWare3,UserInforStruct)
	chap02.GET("/arr",ArrController)
	chap02.GET("/arr_struct",ArrStruct)

	chap02.GET("/map",MapController)
	chap02.GET("/map_struct",MapStruct)

	chap02.GET("/slice",SliceController)

	chap02.GET("/param1/:id",Param1)
	chap02.GET("/param2/*id",Param2)

	chap02.GET("/query",GetQueryData)

	chap02.GET("/query_arr",GetQueryArrData)
	chap02.GET("/query_map",GetQueryMapData)

	chap02.GET("/to_user_add",ToUserAdd)
	chap02.POST("/do_user_add",DoUserAdd)

	chap02.GET("/to_user_add2",ToUserAdd2)
	chap02.POST("/do_user_add2",DoUserAdd2)

	chap02.GET("/to_user_add3",ToUserAdd3)
	chap02.POST("/do_user_add3",DoUserAdd3)

	chap02.GET("/to_user_add4",ToUserAdd4)
	chap02.POST("/do_user_add4",DoUserAdd4)

	chap02.GET("/test_to_upload1",ToUpload1)
	chap02.GET("/test_to_upload2",ToUpload2)
	chap02.POST("/test_do_upload1",DoUpload1)
	chap02.POST("/test_do_upload2",DoUpload2)

	chap02.GET("/test_to_upload3",ToUploadFile3)
	chap02.POST("/test_do_upload3",DoUploadFile3)

	chap02.GET("/test_to_upload4",ToUploadFile4)
	chap02.POST("/test_do_upload4",DoUploadFile4)

	chap02.GET("/output_json",OutJson)
	chap02.GET("/output_asciijson",OutAsciiJson)
	chap02.GET("/output_jsonp",OutJsonp)
	chap02.GET("/output_oure_json",OutPureJsonp)
	chap02.GET("/output_secure_json",OutSecureJsonp)
	chap02.GET("/output_xml",OutXml)
	chap02.GET("/output_yaml",OutYaml)
	chap02.GET("/output_proto",OutProto)

	chap02.GET("/redirect_a",RedirectA)
	chap02.GET("/redirect_b",RedirectB)

}
