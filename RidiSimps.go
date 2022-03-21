package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

func loginData() (id string, password string) {
	fmt.Print("ID : ")
	fmt.Scanln(&id)
	fmt.Print("PASSWORD : ")
	fmt.Scanln(&password)

	return id, password
}

func login(id string, password string) (*http.Response, error) {

	// 로그인 request의 body를 작성하는 과정.
	// POST method, content-type : multipart/form-data

	loginUri := "https://ridibooks.com/account/action/login"

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

	return http.Post(loginUri, "multipart/form-data", &body)
}

func main() {
	id, password := loginData()

	resp, err := login(id, password)
	if err != nil {
		fmt.Println(err, ": login failed")
	}

}
