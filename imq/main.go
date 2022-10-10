package imq

import (
	"fmt"
	"time"
)

func main() {
	OnceTopic()
}

var topic = fmt.Sprintf("Golang梦工厂")

// 一个topic 测试
func OnceTopic() {
	m := NewClient()
	m.SetConditions(10)

	ch, err := m.Subscribe(topic)
	if err != nil {
		fmt.Println("subscribe failed")
		return
	}
	go OncePub(m)
	OnceSub(ch, m)
	defer m.Close()
}

// 定时推送
func OncePub(c *Client) {
	t := time.NewTicker(10 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			err := c.Publish(topic, "asong is butty")
			if err != nil {
				fmt.Println("pub message failed")
			}
		default:

		}
	}
}

// 接受订阅消息
func OnceSub(m <-chan interface{}, c *Client) {
	for {
		val := c.GetPayLoad(m)
		fmt.Printf("get message is %s\n", val)
	}
}
