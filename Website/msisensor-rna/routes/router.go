package routes

import (
	v1 "msisensor-rna/api/v1"
	"msisensor-rna/middleware"
	"msisensor-rna/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "./web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.StaticFile("/favicon.png", "./web/front/dist/favicon.png")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 样本模块的路由接口
		auth.POST("sample/add", v1.AddSample)
		auth.DELETE("sample/:sid", v1.DeleteSample)
		auth.GET("samples", v1.GetSamples)
		auth.PUT("sample/:sid", v1.EditSample)
		auth.GET("detection/:sid", v1.BulkDetection)

		// 上传样本
		auth.POST("sample/upload", v1.Upload)

	}
	router := r.Group("api/v1")

	{
		router.POST("user/add", v1.AddUser)
		// router.GET("users", v1.GetUsers)

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

	}

	r.Run(utils.HttpPort)
}
