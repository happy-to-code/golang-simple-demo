package main

func main() {
	bc := NewBlockChain("13K6dbnAeu5AdSqFhYfigVhoTN8zQrq4LSMbhrka")
	cli := CLI{bc}
	cli.Run()
}
