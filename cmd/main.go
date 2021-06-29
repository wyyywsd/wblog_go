package main

import (
	"net/http"
	"wblogApi/api/handler"
	"wblogApi/internal/db"
	"wblogApi/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//数据库初始化
	db.InitDb()
	//自动更新数据库字段信息
	db.W_Db.AutoMigrate(&models.User{}, &models.Article{})
	r := gin.Default()
	//设置对外访问的静态文件
	r.Static("/public/articles", "../public/articles")
	r.Static("/public/user_profile", "../public/user_profile")
	//注册一个路由组 里面都是 给vue前端准备的api
	apiGroup := r.Group("/api/v1")
	{
		//对外获取token的api  暂时也用作登陆api
		apiGroup.POST("/auth", handler.AuthHandler)
		//注册api
		apiGroup.POST("/signup", handler.SignUp)

		apiGroup.GET("/home", handler.JWTAuthMiddleware(), homeHandler)
		apiGroup.POST("/create_article", handler.JWTAuthMiddleware(), handler.CreateArticleHandler)
		apiGroup.GET("/show_article/:article_id", handler.JWTAuthMiddleware(), handler.ShowArticle)
		apiGroup.GET("/show_article_list", handler.JWTAuthMiddleware(), handler.ShowArticleList)
		apiGroup.POST("/upload", handler.UploadArticlePhoto)
		apiGroup.POST("/uploadUserProfilePhoto", handler.UploadUserProfilePhoto)
		apiGroup.POST("/update_user", handler.JWTAuthMiddlewareA(), handler.UpdateUserInfo)
		apiGroup.GET("/demo", handler.JWTAuthMiddleware(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "登陆成功",
				"code":    2000,
			})
		})
	}

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8080")
}
func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}
