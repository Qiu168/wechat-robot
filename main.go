package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

func main() {
	//print(getXinghuo("你好"))
	//fmt.Scanf("t")
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() {
			if strings.Contains(msg.Content, "@湫 ") {
				//user, _ := msg.SenderInGroup()
				resp := getXinghuo(msg.Content[len("@湫 "):])
				fmt.Println(resp)
				text, err := msg.ReplyText(resp)
				if err != nil {
					fmt.Errorf("%v", err)
				}
				fmt.Println(text)
			}
		}
	}
	//bot.MessageHandler = func(msg *openwechat.Message) {
	//	if msg.IsText() {
	//		if strings.Contains(msg.Content, "湫") {
	//			println(msg.Content)
	//			msg.ReplyText("666")
	//		}
	//
	//	}
	//}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	_, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	//friends, err := self.Friends()
	//fmt.Println(friends, err)

	//// 获取所有的群组
	//groups, err := self.Groups()
	//fmt.Println(groups, err)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
