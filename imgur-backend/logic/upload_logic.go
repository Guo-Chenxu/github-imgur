package logic

import (
	"encoding/json"
	"fmt"
	"imgur-backend/conf"
	"io"
	"net/http"
	"strings"
)

// Upload 发送请求
func Upload(message, filename, content string) error {
	// 准备参数并发送请求
	url := fmt.Sprintf("https://api.github.com/repos/%s/imgs/contents/%s/%s", conf.Conf.User, conf.Conf.Path, filename)
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
		return err
	}
	req.Header.Add("Accept", conf.Conf.Accept)
	req.Header.Add("Authorization", "token "+conf.Conf.Authorization)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Host", "api.github.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// 对返回值校验, 正常的返回值中会带有文件名
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	ok := strings.Contains(string(resp), filename)
	if !ok {
		return fmt.Errorf(string(resp))
	}
	return nil
}
