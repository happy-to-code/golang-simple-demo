package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Gender int    `json:"gender,omitempty"`
}

func main() {
	var u User
	u = User{
		Name:   "XiaoMing",
		Age:    15,
		Gender: 1,
	}
	userBytes, err := json.Marshal(u)
	if err != nil {
		errors.Wrap(err, "JSON序列化失败")
		return
	}
	fmt.Println(string(userBytes))

	var s = `"name":"XiaoMing","age":15,"gender":1}`
	err = json.Unmarshal([]byte(s), &u)
	if err != nil {
		fmt.Printf("err0:%+v\n", err)
		err = errors.Wrap(err, "JSON反序列化失败")
		fmt.Printf("err1:%+v\n", err)
		fmt.Println(fmt.Errorf("err2:%w\n", err).Error())
		fmt.Printf("err3:%s\n", err.Error())

	}
	fmt.Println(u)

}
