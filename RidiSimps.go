package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
)

func loginData() (id string, password string) {
	fmt.Print("ID : ")
	fmt.Scanln(&id)
	fmt.Print("PASSWORD : ")
	fmt.Scanln(&password)

	return id, password
}

func login(id string, password string) {

	// 로그인 request의 body를 작성하는 과정.
	// POST method, content-type : multipart/form-data

	fieldNames := []string{ // 리디북스 로그인 request의 body field들.
		"user_id",
		"password",
		"cmd",
		"return_url",
		"return_query_string",
		"device_id",
		"msg"}

	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	w.WriteField(fieldNames[0], id)
	w.WriteField(fieldNames[1], password)
	w.WriteField(fieldNames[2], "login")

	for _, fn := range fieldNames[3:] {
		w.CreateFormField(fn)
	}

	w.Close()

}

func main() {
	var id string
	var password string

	id, password = loginData()
	login(id, password)

}
