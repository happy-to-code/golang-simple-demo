package main

import "fmt"
import "github.com/google/wire"

// Message 消息
type Message struct {
	msg string
}

// Greeter 打招呼
type Greeter struct {
	Message Message
}

// Event 事件
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func (g Greeter) Greet() Message {
	return g.Message
}

// 使用wire前
func main() {
	message := NewMessage("hello world")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
	fmt.Println("-----------------------")

	event2 := InitializeEvent2("hello_world")

	event2.Start()
}

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage(msg))
	return Event{} // 返回值没有实际意义，只需符合函数签名即可
}

func InitializeEvent2(msg string) Event {
	message := NewMessage(msg)
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	return event
}
