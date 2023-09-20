package controller

import (
	"imgur-backend/conf"
	"imgur-backend/logic"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Image struct {
	Name string `json:"name"`
	File string `json:"file"`
}

// Upload 上传图片
func Upload(ctx *gin.Context) {
	// 获取参数
	var image Image
	err := ctx.ShouldBind(&image)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err.Error()})
		return
	}

	// 准备请求的参数并发送请求
	content := image.File
	postfix := strings.Split(image.Name, ".")
	filename := time.Now().Format("200601021504050") + "." + postfix[len(postfix)-1]

	url, err := "", nil
	switch conf.Conf.Bed {
	case "github":
		message := conf.Conf.GithubConfig.Message
		url, err = logic.UploadByGithub(message, filename, content)
	case "gitee":
		message := conf.Conf.GiteeConfig.Message
		url, err = logic.UploadByGitee(message, filename, content)

	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"url":    url,
	})
}
