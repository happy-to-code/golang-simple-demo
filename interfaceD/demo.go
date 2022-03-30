package main

import "fmt"

type USB interface {
	working()
	stop()
}

// 定义Phone对象
type Phone struct {
}

// 定义Cameron对象
type Cameron struct {
}

// 定义Computer对象
type Computer struct {
}

// 分别让两个对象实现USB的方法
func (p Phone) working() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) stop() {
	fmt.Println("手机停止工作")
}
func (c Cameron) working() {
	fmt.Println("照相机开始工作。。。")
}
func (c Cameron) stop() {
	fmt.Println("照相机停止工作")
}

func (c Computer) work(usb USB) { // 具有相同接口的对象（struct）都可以作为函数参数(包含多态、高内聚低耦合的思想)
	usb.working() // 根据上下文判断是Cameron 还是 phone，实现了多态
	usb.stop()
}

// 自定义类型的变量 都可以实现接口，不一定是struct
type MyInt int

func (i MyInt) working() {
	fmt.Println("自定义类型的  working（） 方法")
}
func (i MyInt) stop() {
	fmt.Println("自定义类型的  stop（） 方法")
}
func main() {
	c := Computer{}
	ca := Cameron{}
	ph := Phone{}
	c.work(ca) // 由于Cameron实现了USB的接口，所以类型可以和USB进行匹配
	c.work(ph) // 由于phone实现了USB的接口，所以类型可以和USB进行匹配
	fmt.Println("------------")
	var usb USB
	var myInt MyInt
	usb = myInt // 赋值给接口（包含方法），myInt需要实现该接口的所有方法。
	usb.stop()
	myInt.working()
	fmt.Println("==================")
	var p Phone
	p.working()
}
