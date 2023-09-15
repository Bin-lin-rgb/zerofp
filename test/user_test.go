package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"zerofp/define"
	"zerofp/pkg"
)

var userServiceAddr = "http://127.0.0.1:13000"

func TestUserLogin(t *testing.T) {
	m := define.M{
		"username": "get",
		"password": "123456",
	}
	data, _ := json.Marshal(m)
	rep, err := pkg.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
