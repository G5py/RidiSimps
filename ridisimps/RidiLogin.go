package ridisimps

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
)

type LoginData struct {
	Id string
	Pw string
}

func Login() *http.Client {
	client := makeClient()
	client.Do(getLoginDataByConsole().toLoginRequest())
	return client
}

func makeClient() *http.Client {
	cookiejar, _ := cookiejar.New(nil)
	return &http.Client{Jar: cookiejar}
}

func getLoginDataByConsole() LoginData {
	var id, pw string
	fmt.Print("ID : ")
	fmt.Scan(&id)
	fmt.Print("Password : ")
	fmt.Scan(&pw)
	return LoginData{id, pw}
}

func (loginData LoginData) toLoginRequest() *http.Request {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	w.WriteField("user_id", loginData.Id)
	w.WriteField("password", loginData.Pw)
	w.WriteField("cmd", "login")
	w.Close()

	request, _ := http.NewRequest("POST", "https://ridibooks.com/account/action/login", &body)
	request.Header.Set("Content-Type", "multipart/form-data; boundary="+w.Boundary())

	return request
}
