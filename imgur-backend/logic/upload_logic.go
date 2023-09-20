package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"imgur-backend/conf"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// Upload 使用github上传
func UploadByGithub(message, filename, content string) (string, error) {
	// 准备参数并发送请求
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s", conf.Conf.GithubConfig.User, conf.Conf.GithubConfig.Repo, conf.Conf.GithubConfig.Path, filename)
	method := "PUT"

	type UploadReq struct {
		Message string `json:"message"`
		Content string `json:"content"`
	}

	uploadReq, _ := json.Marshal(UploadReq{Message: message, Content: content})
	payload := strings.NewReader(string(uploadReq))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", conf.Conf.Accept)
	req.Header.Add("Authorization", "token "+conf.Conf.Authorization)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Host", "api.github.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 对返回值校验, 正常的返回值中会带有文件名
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	ok := strings.Contains(string(resp), filename)
	if !ok {
		return "", fmt.Errorf(string(resp))
	}
	return conf.Conf.CDN + filename, nil
}

// Upload 使用gitee上传
func UploadByGitee(message, filename, content string) (string, error) {
	// 准备参数
	url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s/%s", conf.Conf.GiteeConfig.Owner, conf.Conf.GiteeConfig.Repo, conf.Conf.GiteeConfig.Path, filename)
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("access_token", conf.Conf.GiteeConfig.AccessToken)
	_ = writer.WriteField("content", content)
	_ = writer.WriteField("message", conf.Conf.GiteeConfig.Message)
	if conf.Conf.GiteeConfig.Branch != "" {
		conf.Conf.GiteeConfig.Branch = "master"
	}
	_ = writer.WriteField("branch", conf.Conf.GiteeConfig.Branch)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "gitee.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------325049816210664367164649")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	// 处理返回值
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ok := strings.Contains(string(body), filename)
	if !ok {
		return "", fmt.Errorf(string(body))
	}
	resp := fmt.Sprintf("https://gitee.com/%s/%s/raw/%s/%s/%s", conf.Conf.GiteeConfig.Owner, conf.Conf.GiteeConfig.Repo, conf.Conf.GiteeConfig.Branch, conf.Conf.GiteeConfig.Path, filename)
	return resp, nil
}
