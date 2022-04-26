package main

import (
	"fmt"
	"os/exec"
)

func main() {
	for i := 0; i < 300; i++ {
		//
		// 	// 有GUI调用
		go exec.Command(`cmd`, `/c`, `start`, `https://www.jszzb.gov.cn/col22/81608.html`).Start()
		fmt.Println("------------------------>", i)
	}

	// 无GUI调用
	// for i := 0; i < 500; i++ {
	// 	cmd := exec.Command(`cmd`, `/c`, `start`, `https://www.jszzb.gov.cn/col22/81608.html`)
	// 	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	// 	cmd.Start()
	// }
	for {

	}
}
