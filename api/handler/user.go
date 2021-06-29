package handler

import (
	"encoding/json"
	"fmt"
	"wblogApi/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignIn(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "登陆成功",
		"code":    2000,
	})

}

func SignUp(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)
	//获取到了 前端传来的json数据  user_name pass_word
	//fmt.Println(m["user_name"])
	userName := m["user_name"]
	passWord := m["pass_word"]
	//这里增加校验  如果用户名存在 就注册失败
	result, _ := models.FindUserByUserName(fmt.Sprint(userName))
	if result.Error != nil {
		//err 不为空 说明 查不到这条记录
		if result.Error == gorm.ErrRecordNotFound {
			//到这里说明 数据库里没有这个用户名 这时候 就可以注册了
			models.CreateUser(fmt.Sprint(userName), fmt.Sprint(passWord))
			c.JSON(200, gin.H{
				"message": "注册成功",
				"code":    2000,
			})
		} else {
			//到这里说明  不是查不到记录的报错 , 就先返回有未知问题吧
			c.JSON(200, gin.H{
				"message": "有未知问题",
				"code":    2009,
			})
		}

	} else {
		//err为空 说明查到记录了
		c.JSON(200, gin.H{
			"message": "用户名已存在",
		})
	}

}

func UpdateUserInfo(c *gin.Context) {
	fmt.Println("进入修改用户的action了")
	//获取JWTAuthMiddleware 中间件中保存的用户信息
	userName, _ := c.Get("username")
	//获取用户
	_, user := models.FindUserByUserName(fmt.Sprint(userName))
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)
	//更新用户信息
	result := models.UpdateUserInfo(user, m)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"message":      result.Error,
			"code":         2010,
			"tokenCode":    2000,
			"tokenMessage": "",
		})
	} else {
		//更新成功后 token应该更新

		c.JSON(200, gin.H{
			"message":      "更新成功",
			"code":         2000,
			"tokenCode":    2000,
			"tokenMessage": "",
		})
	}
}
