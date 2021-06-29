package handler

import (
	"fmt"
	"net/http"
	"strings"
	"wblogApi/internal/jwt"
	"wblogApi/internal/models"

	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	//用户发送用户名和密码过来
	var user models.User
	err := c.ShouldBind(&user)
	fmt.Println(user.UserName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    2001,
			"message": "无效的参数",
		})
		return
	}

	// 校验用户名和密码是否正确
	//通过用户名先获取用户信息
	result, checkUser := models.FindUserByUserName(user.UserName)
	fmt.Println(result.Error)
	if result.Error != nil {
		//有错误, 用这个用户名查不到对应的用户数据
		c.JSON(http.StatusOK, gin.H{
			"code":    2002,
			"message": "用户名不存在",
		})
		return
	} else {
		//到这里证明 这个用户名在数据库里是存在的
		//在这里检查密码是否正确
		if checkUser.PassWord == user.PassWord {
			//用户名密码全部正确
			// 生成Token
			tokenString, _ := jwt.GenToken(user.UserName)

			c.JSON(http.StatusOK, gin.H{
				"code":        2000,
				"message":     "登陆成功,欢迎回来!" + checkUser.UserName,
				"data":        gin.H{"token": tokenString},
				"currentUser": checkUser,
			})
			return
		} else {
			//密码不对
			c.JSON(http.StatusOK, gin.H{
				"code":    2003,
				"message": "密码不对,鉴权失败",
			})
			return
		}
	}

}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddlewareA() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"tokenCode":    2004,
				"tokenMessage": "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"tokenCode":    2005,
				"tokenMessage": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"tokenCode":    2006,
				"tokenMessage": "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.UserName)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

// JWTAuthMiddlewareNext 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定

		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.Set("tokenCode", 2004)
			c.Set("tokenMessage", "请求头中auth为空")
			c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		} else {

			// 按空格分割
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				c.Set("tokenCode", 2005)
				c.Set("tokenMessage", "请求头中auth格式有误")
				c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
			} else {
				// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
				mc, err := jwt.ParseToken(parts[1])
				if err != nil {
					c.Set("tokenCode", 2006)
					c.Set("tokenMessage", "无效的Token")
					c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
				} else {
					c.Set("tokenCode", 2000)
					c.Set("tokenMessage", "认证通过")
					// 将当前请求的username信息保存到请求的上下文c上
					c.Set("username", mc.UserName)
					c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
				}
			}
		}

	}
}
