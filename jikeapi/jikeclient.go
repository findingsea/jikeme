package jikeclient

import (
	"github.com/mdp/qrterminal"
	"os"
	"log"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
)

const host = "app.jike.ruguoapp.com"
var httpClient = &http.Client{}
var jikeCreateSession = fmt.Sprintf("http://%s/sessions.create", host)
var jikeWait4Login = fmt.Sprintf("https://%s/sessions.wait_for_login", host)
var jikeWait4Confirm = fmt.Sprintf("https://%s/sessions.wait_for_login", host)
var headers = map[string]string{
	"Origin":          "https://web.okjike.com",
	"Referer":         "https://web.okjike.com",
	"User-Agent":      "jikeme",
	"Accept":          "application/json",
	"Content-Type":    "application/json",
	"Accept-Language": "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7",
	"App-Version":     "5.3.0",
	"platform":        "web",
}

// JikeClient 即刻 API 客户端
type JikeClient struct {
	Token string `string:"Token"`
}

func client() *http.Client {
	return httpClient
}

func (j *JikeClient) newRequst(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal("Get new request fail")
	}
	if j.Token != "" {
		req.Header.Set("x-jike-app-auth-jwt", j.Token)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req
}

// Login 登录
func (j *JikeClient) Login() {
	uuid, err := j.CreateSession()
	if err != nil {
		log.Fatal("Create session err: ", err)
	}
	log.Println("Get uuid: ", uuid)
	generateQRCode(uuid)
}

// CreateSession 创建 session，获取 uuid
func (j *JikeClient) CreateSession() (string, error) {
	req := j.newRequst("GET", jikeCreateSession, nil)
	var data struct {
		UUID string `json:"uuid"`
	}
	resp, err := client().Do(req)
	if err != nil {
		return "", err
	}
	log.Println("get resp succes")
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(raw))
	err = json.Unmarshal(raw, &data)
	if err != nil {
		return "", err
	}
	return data.UUID, nil
}

// generateQRCode 生成二维码
func generateQRCode(uuid string) {
	url := `jike://page.jk/web?url=https%3A%2F%2Fruguoapp.com%2Faccount%2Fscan%3Fuuid%3D` + uuid + `&displayHeader=false&displayFooter=false`
	qrterminal.GenerateHalfBlock(url, qrterminal.L, os.Stdout)
}
