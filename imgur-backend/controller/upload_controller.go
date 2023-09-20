package controller

import (
	"encoding/base64"
	"imgur-backend/conf"
	"imgur-backend/logic"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Upload 上传图片
func Upload(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("name")
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err.Error()})
		return
	}

	// 将照片转换成 base64 编码
	file, _ := fileHeader.Open()
	fileData := make([]byte, fileHeader.Size)
	_, err = file.Read(fileData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err.Error()})
		return
	}

	// 准备请求的参数并发送请求
	content := base64.StdEncoding.EncodeToString(fileData)
	postfix := strings.Split(name, ".")
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
