package publishsubscribe

import (
	"sync"
	"time"
)

// 一个订阅者的类型是一个channel
type subscriber chan interface{}

type topicFunc func(v interface{}) bool

// 一个发布者，它会记录当前所有的订阅者

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

// Pulisher的new方法，可以设置发布超时和缓存队列长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
		// 这里没有设置读写锁，其实也是不用的
		// 另外，一定要记住，锁的本质是锁goroutine而不是某个变量。
	}
}

// 添加一个新的订阅者，它会订阅一个主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 封装一下上面那个函数，订阅一个空的Topic
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 退出订阅，干两件事
// 1. 将自己从订阅者列表中拿走
// 2. close这个channel
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// 全部取消订阅
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 发布主题
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// 这里有一个很关键的函数，就是topic其实会返回一个bool来看是否发布的和我订阅的内容一致。
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// 向所有的订阅者发布主题参数v，经过上面的sendTopic函数，我们可以将正确的消息发送到正确的订阅者上面。
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
}
