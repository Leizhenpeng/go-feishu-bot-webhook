package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CatchZeng/feishu/pkg/feishu"
)

func sendDrinkMessage(client *feishu.Client) {
	msg := feishu.NewInteractiveMessage()

	t := time.Now()
	tf := t.Format("2006-01-02 15:04:05")
	msg.MsgType = "interactive"

	const card string = `{
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "🚒 喝水提醒 ~ %s ~"
    },
    "template": "wathet"
  },
  "elements": [
    {
      "alt": {
        "content": "",
        "tag": "plain_text"
      },
      "img_key": "img_v2_a0d2c3e6-5b48-4637-bb7d-fe1b30694a8g",
      "tag": "img",
      "mode": "crop_center",
      "compact_width": false
    },
    {
      "tag": "div",
      "text": {
        "content": "朋友们，记得喝水哦！喝水有助于保持健康、清醒、高效。站起来喝一杯吧！",
        "tag": "lark_md"
      }
    },
    {
      "tag": "markdown",
      "content": "\n查看 [喝水秘籍](https://www.bilibili.com/video/BV1Gh411p72E/?spm_id_from=333.337.search-card.all.click) <at id=all></at>"
    }
  ]
}`
	msg.Card = fmt.Sprintf(card, tf)

	_, resp, err := client.Send(msg)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(resp)
}

func sendPlainMessage(client *feishu.Client) {
	msg := feishu.NewTextMessage()
	msg.Content.Text = "Hello, world!"
	_, resp, err := client.Send(msg)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(resp)
}

func main() {
	token := "xxxx"
	secret := "xxxx"

	client := feishu.NewClient(token, secret)

	sendPlainMessage(client)
	sendDrinkMessage(client)

	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			sendDrinkMessage(client)
		}
	}

}
