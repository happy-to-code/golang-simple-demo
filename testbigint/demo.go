package main

import (
	"fmt"
	"math/big"
)

func main() {
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	temInt := big.Int{}
	fmt.Println("111:", temInt)
	fmt.Println("222:", &temInt)
	setString, b := temInt.SetString(targetStr, 16)
	fmt.Println("333:", temInt)
	fmt.Println("444:", &temInt)
	fmt.Println("-----------------")
	fmt.Println(b)
	fmt.Println(setString)
	fmt.Println(&setString)

}
