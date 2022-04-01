package main

import "fmt"

func main() {
	/*
		%d          十进制整数
		%x, %o, %b  十六进制，八进制，二进制整数。
		%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
		%t          布尔：true或false
		%c          字符（rune） (Unicode码点)
		%s          字符串
		%q          带双引号的字符串"abc"或带单引号的字符'c'
		%v          变量的自然形式（natural format）
		%T          变量的类型
		%%          字面上的百分号标志（无操作数）

	*/
	fmt.Printf("%q\n", "abc123")
	fmt.Printf("%q\n", "a")
	fmt.Printf("%q\n", "2")
	fmt.Printf("%q\n", "12")
	fmt.Printf("%f\n", 3.141592653589793)
	fmt.Printf("%g\n", 3.141592653589793)
	fmt.Printf("%e\n", 3.141592653589793)
	fmt.Printf("%%\n")
}
