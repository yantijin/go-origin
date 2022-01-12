package main

import (
	"fmt"
	"originyan/mail"
	"originyan/webhook"
	"strconv"
	"time"
)


func sendmail() {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！
		</body>
		</html>`

	mailType := "html"//发送邮件的人名称
	u := mail.User{}
	err := u.ReadConfig("config.json")
	if err != nil {
		fmt.Println(err)
	} else {
		u.Body = body
		u.Mailtype = mailType
	}
	//fmt.Println(u)
	err = u.SendToMail()
	if err != nil {
		fmt.Println("发送失败")
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功！")
	}

}

func send_msg() {
	a := webhook.WebHook{}
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	err := a.ReadConfig("./config1.json")
	if err != nil {
		fmt.Println("解析失败")
	}
	a.CreateDingJsonMsg("测试", "测试go发送webhook消息", []string{"18811532609"}, []string{"闫媞锦"})
	fmt.Println(a)
	a.SendDingMsg(timestamp)
}

func send_qywx_msg() {
	a := webhook.WebHook{}
	err := a.ReadConfig("./config2.json")
	if err != nil {
		fmt.Println("解析失败")
	}
	a.CreateQywxMsg("测试go发送webhook消息")
	fmt.Println(a)
	a.SendQywxMsg()
}
