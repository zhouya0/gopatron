## 发布订阅模式

发布订阅模式的设计原则是将那些从观察者或者发布者获得消息的人的解耦，意味着不需要编程直接发送给特定接受者消息。


为了实现这一点，中介称为"消息代理"或者"事件总线"，接收已发布的消息，然后将它们路由到订阅服务器。