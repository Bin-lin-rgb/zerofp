package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"zerofp/pkg"
)

var adminServiceAddr = "http://localhost:13001"

func TestProductCreate(t *testing.T) {
	m2 := map[string]string{
		"name": "name112",
		"desc": "desc2",
	}
	data, _ := json.Marshal(m2)

	rep, err := pkg.HttpPost(adminServiceAddr+"/product/create", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductDelete(t *testing.T) {
	m := map[string]string{
		"identity": "3501dc56-b080-40bd-8e1d-b5dc69b52609",
	}
	data, _ := json.Marshal(m)

	rep, err := pkg.HttpDelete(adminServiceAddr+"/product/delete", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductModify(t *testing.T) {
	m := map[string]string{
		"identity": "6330a8e8-45d9-4c4f-aeb7-ef9373e5525d",
		"name":     "-------",
		"desc":     "-------",
	}
	data, _ := json.Marshal(m)

	rep, err := pkg.HttpPut(adminServiceAddr+"/product/modify", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductList(t *testing.T) {

	rep, err := pkg.HttpGet(adminServiceAddr + "/product/list?page=1&size=20&name=")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
