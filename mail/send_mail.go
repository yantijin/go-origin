package mail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
) //TODO: 发送附件功能，cc功能

type User struct {
	User string `json:"user"`
	SendUserName string `json:"sendUserName"`
	Password string `json:"password"`
	Host string `json:"host"`
	To string `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Mailtype string `json:"mailtype"`// html或者其他内容
}

func (r *User)SendToMail() error {
	hp := strings.Split(r.Host, ":")
	auth := smtp.PlainAuth("", r.User, r.Password, hp[0])
	var content_type string
	if r.Mailtype == "html" {
		content_type = "Content-Type: text/" + r.Mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + r.To + "\r\nFrom: " + r.SendUserName+"<"+r.User+">" + "\r\nSubject: " +r.Subject+ "\r\n" + content_type + "\r\n\r\n" + r.Body)
	send_to := strings.Split(r.To, ";")
	err := smtp.SendMail(r.Host, auth, r.User, send_to, msg)
	return err
}

func (r *User) ReadConfig(path string) (err error) {
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

