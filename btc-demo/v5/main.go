package main

// import "fmt"

func main() {
	bc := NewBlockChain("1KWVA6Qbp4afppnQ3MNQTa9NanM7gwjCqt")
	cli := CLI{bc}
	cli.Run()
}
