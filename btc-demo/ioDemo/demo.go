package main

import (
	"io/ioutil"
	"log"
)

func main() {
	var data = `{"name":"xiaoming","age":17}`
	err := ioutil.WriteFile("1.json", []byte(data), 0600)
	if err != nil {
		log.Panicln(err)
	}

}
