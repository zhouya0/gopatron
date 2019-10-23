package publishsubscribe

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPublishSubscribe(t *testing.T) {
	fmt.Println("test")
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()
	all := p.Subscribe()
	// all是一个订阅者 它要求的topic是空的，也就是全部都需要。所以
	// publish的所有内容它都会收到。
	golang := p.SubscribeTopic(func(v interface{}) bool { //golang也是个订阅者，它要求消息包含“golang”
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world!")
	p.Publish("hello, golang")

	go func() {
		for msg := range all {
			fmt.Println("golang", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang", msg)
			if msg != "hello, golang" {
				t.Fatal("test failed")
			}
		}
	}()
	time.Sleep(3 * time.Second)
}

// 我们回顾一下这里的发布订阅模型：
// 1. 订阅：订阅之后是将一个channel和topic加入到了publisher的subscribers的map中
// 可以理解为，订阅完成的动作是：我是谁，我要什么样的消息。
// 2. 发布： publisher遍历自己的subscribers然后将消息放进去。
// 其中，送消息的时候会过一个send检查函数，表示topic不对的函数就直接返回，符合的topic就把消息塞进channel里。
