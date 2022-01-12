package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

type WebHook struct {
	Webhook string `json:"webhook"`
	Msg string
	AppSecret string `json:"app_secret"`
}
type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
}
type Msg struct {
	Title string `json:"title"`
	Text string `json:"text"`
}
type QywxContent struct {
	Content string `json:"content"`
}

type Msg1 struct {
	MsgType string `json:"msgtype"`
	MarkDown Msg `json:"markdown"`
	At At `json:"at"`
}

type QywxMsg struct {
	MsgType string `json:"msgtype"`
	MarkDown QywxContent `json:"markdown"`
}

func getSign(timestamp string, appsecret string) string{
	message := []byte(timestamp + "\n" + appsecret)
	secret := []byte (appsecret)
	hash := hmac.New(sha256.New, secret)
	hash.Write(message)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func (w *WebHook)SendDingMsg(timestamp string) {
	sign := getSign(timestamp, w.AppSecret)
	header := "application/json; charset=utf-8"
	url := w.Webhook + "&timestamp=" + timestamp + "&sign=" + url2.QueryEscape(sign)
	resp, err := http.Post(url, header, strings.NewReader(w.Msg))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func (w *WebHook)SendQywxMsg() {
	header := "application/json; charset=utf-8"
	resp, err := http.Post(w.Webhook, header, strings.NewReader(w.Msg))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func (r *WebHook) ReadConfig(path string) (err error) {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件失败，请检查路径")
		return
	}

	err = json.Unmarshal(jsonFile, &r)
	if err != nil {
		fmt.Println("解析数据失败，请检查")
		return
	}

	return
}

func (r *WebHook)CreateDingJsonMsg(title string, text string, mobiles []string, userIds []string) {
	msg := &Msg1{
		MsgType: "markdown",
		MarkDown: Msg{Title: title, Text: text},
		At: At{AtMobiles: mobiles, AtUserIds: userIds},
	}
	res, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("解析失败")
	}
	r.Msg = string(res)
}

func (r *WebHook)CreateQywxMsg(text string) {
	msg := QywxMsg{
		MsgType: "markdown",
		MarkDown: QywxContent{Content: text},
	}
	res, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("解析失败")
	}
	r.Msg = string(res)
}
