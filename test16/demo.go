package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"os/exec"
	"time"
)

func main() {
	exec.Command(`cmd`, `/c`, `start`, `https://www.jszzb.gov.cn/col22/81608.html`).Start()
	// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
	// robotgo.MoveMouse(90, 50)
	robotgo.MoveMouse(90, 50)
	// 向上滚动：3行
	// robotgo.ScrollMouse(3, `up`)
	// // 向下滚动：2行
	// robotgo.ScrollMouse(2, `down`)

	// 按下鼠标左键
	// 第1个参数：left(左键) / center(中键，即：滚轮) / right(右键)
	for i := 0; i < 2200; i++ {

		robotgo.MouseClick(`left`, false)
		time.Sleep(time.Millisecond * 80)
	}
	// 第2个参数：是否双击

	// 按住鼠标左键
	// robotgo.MouseToggle(`down`, `left`)
	// 解除按住鼠标左键
	// robotgo.MouseToggle(`up`, `left`)

	fmt.Println("===========================")
}
