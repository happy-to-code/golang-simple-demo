package main

import "fmt"

type Reader interface {
	Reader()
}

type Writer interface {
	Write()
}

type ReadAndWrite interface {
	Reader
	Writer
}

type Java struct{}
type Golang struct{}

func (j *Java) Reader() {
	fmt.Println("i am java reader")
}
func (j *Java) Write() {
	fmt.Println("i am java writer")
}

func (j *Golang) Reader() {
	fmt.Println("i am golang reader")
}
func (j *Golang) Write() {
	fmt.Println("i am golang writer")
}

type User struct {
	rw   ReadAndWrite
	Name string
	Age  int
}

func main() {
	u := &User{
		rw:   ReadAndWrite(&Golang{}),
		Name: "dx",
		Age:  10,
	}
	u.rw.Write()
	// var rw ReadAndWrite
	// rw = &Java{}
	// rw.Write()
	// rw.Reader()
	// rw = &Golang{}
	// rw.Write()
	// rw.Reader()
	// write(&Java{})
	// read(&Java{})
}

type fun ReadAndWrite

func write(t ReadAndWrite) {
	fun.Write(t)
}
func read(t ReadAndWrite) {
	fun.Reader(t)
}
