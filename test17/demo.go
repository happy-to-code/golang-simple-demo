package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"os/exec"
)

func main() {
	exec.Command(`cmd`, `/c`, `start`, `https://www.jszzb.gov.cn/col22/81608.html`).Start()
	// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
	robotgo.MoveMouseSmooth(588, 50)
	// robotgo.MoveMouse(588, 64)
	// 向上滚动：3行
	// robotgo.ScrollMouse(3, `up`)
	// // 向下滚动：2行
	// robotgo.ScrollMouse(2, `down`)

	// 按下鼠标左键
	// 第1个参数：left(左键) / center(中键，即：滚轮) / right(右键)
	for i := 0; i < 500; i++ {

		robotgo.MouseClick(`left`, false)
		// 一直按住 A键不放
		robotgo.KeyToggle(`enter`, `down`)
		// 解除按住 A键
		robotgo.KeyToggle(`enter`, `up`)
	}
	// 第2个参数：是否双击

	// 按住鼠标左键
	// robotgo.MouseToggle(`down`, `left`)
	// 解除按住鼠标左键
	// robotgo.MouseToggle(`up`, `left`)

	fmt.Println("===========================")
}
