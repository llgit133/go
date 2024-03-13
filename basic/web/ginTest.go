package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 6.中间件 拦截器
func myHander() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.Set("usersession", "userid-----1")
		context.Next() // 放行

		// if else 判断放行问题
		//context.Abort()//阻止
	}
}

func main() {
	ginServer := gin.Default()
	ginServer.LoadHTMLGlob("basic/web/templates/*")
	ginServer.StaticFile("static", "/basic/web/static")

	// 1.url传参数  url?userId=123&userName=234
	// 2.resultFul传参数 /user/123/liao

	ginServer.GET("/index", func(context *gin.Context) {
		userId := context.Query("userId")
		userName := context.Query("userName")

		context.HTML(http.StatusOK, "index.html", gin.H{"userId": userId, "userName": userName})
		context.JSONP(200, "hello,world")
	})

	ginServer.GET("/user/info/:userId/:userName", func(context *gin.Context) {
		userId := context.Param("userId")
		userName := context.Param("userName")
		context.HTML(http.StatusOK, "index.html", gin.H{"userId": userId, "userName": userName})
		context.JSONP(200, "hello,world")

	})

	// 3.前段给后端json
	ginServer.GET("/json", func(context *gin.Context) {

		// request body
		// [] type
		//data, err := context.GetRawData()
		data, _ := context.GetRawData()

		var m map[string]interface{}

		//包装为json数据 [] type - > json
		_ = json.Unmarshal(data, &m)
		context.JSONP(http.StatusOK, m)

	})

	// 4.表单传参数
	ginServer.POST("/user/add", myHander(), func(context *gin.Context) {

		// 获取用户session
		userSession := context.MustGet("usersession").(string)
		log.Println("=========================>>>" + userSession)
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSONP(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})

	})

	// 5.路由
	ginServer.GET("/testBaidu", func(context *gin.Context) {
		// 重定向 301
		context.Redirect(301, "http://www.baidu.com")
	})

	ginServer.NoRoute(func(context *gin.Context) {
		// 404
		context.HTML(404, "404.html", nil)
		//context.Redirect(404, "")
	})

	// 路由组 /user/add
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.POST("/login")
		userGroup.POST("/logout")
	}

	orderGroup := ginServer.Group("order")
	{
		orderGroup.GET("/add")
		orderGroup.DELETE("/delete")
	}

	ginServer.Run(":8089")

	//ginServer.POST("/user", func(context *gin.Context) {
	//	context.JSONP(200, "hello,world --POST")
	//
	//})
	//
	//ginServer.PUT("/user", func(context *gin.Context) {
	//	context.JSONP(200, "hello,world --PUT")
	//
	//})
	//
	//ginServer.DELETE("/user", func(context *gin.Context) {
	//	context.JSONP(200, "hello,world --DELETE")
	//
	//})

}
