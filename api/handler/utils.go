package handler

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func UploadArticlePhoto(c *gin.Context) {
	fmt.Println("竟然进来了")
	// 单个文件
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println(file.Filename)
	fileName := file.Filename
	dst := fmt.Sprintf("../public/articles/%s", fileName)
	fmt.Println(fileName)
	// 上传文件到指定的目录
	err1 := c.SaveUploadedFile(file, dst)
	fmt.Println(err1)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' localhost:8080/public/", file.Filename),
		"url":     "/public/articles/" + fileName,
	})
}

func UploadUserProfilePhoto(c *gin.Context) {
	fmt.Println("上传了用户的头像")
	file, err := c.FormFile("profilePhoto")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// 注意：下面为了方便，暂时忽略了错误处理
	userId := c.PostForm("ID")
	fileName := userId + path.Ext(file.Filename)
	dst := fmt.Sprintf("../public/user_profile/%s", fileName)
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     "/public/user_profile/" + fileName,
	})

}
