package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"wblogApi/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateArticleHandler(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)
	//获取到了 前端传来的json数据
	articleTitle := m["article_title"]
	article_content := m["article_content"]
	//获取JWTAuthMiddleware 中间件中保存的用户信息
	userName, _ := c.Get("username")
	//获取c 上下文中 储存的token中间件判断信息
	tokenCode, _ := c.Get("tokenCode")
	tokenMessage, _ := c.Get(("tokenMessage"))
	_, user := models.FindUserByUserName(fmt.Sprint(userName))
	result := models.CreateArticle(fmt.Sprint(articleTitle), fmt.Sprint(article_content), user.ID)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"message":      result.Error,
			"code":         2007,
			"tokenCode":    tokenCode,
			"tokenMessage": tokenMessage,
		})
	} else {
		c.JSON(200, gin.H{
			"message":      "保存成功",
			"code":         2000,
			"tokenCode":    tokenCode,
			"tokenMessage": tokenMessage,
		})

	}
}

func ShowArticle(c *gin.Context) {
	//获取c 上下文中 储存的token中间件判断信息
	tokenCode, _ := c.Get("tokenCode")
	tokenMessage, _ := c.Get(("tokenMessage"))
	articleId := c.Param("article_id")
	fmt.Println(articleId)
	article, result := models.FindArticleById(fmt.Sprint(articleId))
	if result.Error != nil {
		c.JSON(200, gin.H{
			"message":      result.Error,
			"code":         2008,
			"tokenCode":    tokenCode,
			"tokenMessage": tokenMessage,
		})
	} else {
		//获取文章作者的信息
		articleUser := models.FindUserById(article.UserId)
		c.JSON(200, gin.H{
			"message":       "success!",
			"code":          2000,
			"article":       article,
			"articleAuthor": articleUser,
			"tokenCode":     tokenCode,
			"tokenMessage":  tokenMessage,
		})

	}
}

//http://192.168.50.200:8081/api/v1/show_article_list?pageSize=2&currentPage=1
func ShowArticleList(c *gin.Context) {
	//获取c 上下文中 储存的token中间件判断信息
	tokenCode, _ := c.Get("tokenCode")
	tokenMessage, _ := c.Get(("tokenMessage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	articles, result, articleTotal := models.FindAllArticles(pageSize, currentPage)

	if result.Error != nil {
		c.JSON(200, gin.H{
			"message":      "文章列表获取失败,请重新获取!",
			"code":         2009,
			"tokenCode":    tokenCode,
			"tokenMessage": tokenMessage,
		})
	} else {
		// 创建一个 保存 map[string]interface{} 的数组 用于保存 下面的articleDetail  返回给客户端
		var articleList [](map[string]interface{})
		//创建一个 map  用于保存  用户 和文章的信息
		var articleDetail = map[string]interface{}{}
		//遍历获取到的文章数据   将文章 和文章对应的作者  添加到 articleList
		for _, article := range articles {
			articleUser := models.FindUserById(article.UserId)
			articleDetail["article"] = article
			articleDetail["articleAuthor"] = articleUser
			articleList = append(articleList, articleDetail)
			articleDetail = map[string]interface{}{}
		}
		c.JSON(200, gin.H{
			"message":      "success!",
			"tokenCode":    tokenCode,
			"tokenMessage": tokenMessage,
			"code":         2000,
			"articleList":  articleList,
			"articleTotal": articleTotal,
		})
	}
}
