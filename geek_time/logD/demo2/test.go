package main

import (
	log "github.com/sirupsen/logrus"
	l "log"
)

func main() {
	log.Print("call Print: line1")
	log.Println("call Println: line2")

	l.Print("call Print: line1")
	l.Println("call Println: line2")
}
