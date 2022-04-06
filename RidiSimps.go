package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// 콘솔창으로 id와 비밀번호를 입력받습니다.
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

	// 사용자로부터 id와 비밀번호를 입력받음.
	id, password := loginData()

	// id와 비밀번호로 ridibooks에 로그인.
	loginResp, err := login(id, password)
	if err != nil {
		fmt.Println(err, ": login failed")
		panic(err)
	} else {
		fmt.Println("로그인 성공")
	}

	// 반복적인 get 호출을 위해 client 선언.
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err, ": cookiejar.new error")
	}
	client := &http.Client{Jar: jar}

	ridiUrl, err := url.Parse("https://ridibooks.com")
	if err != nil {
		fmt.Println(err, ": url parse error")
		panic(err)
	}

	// 로그인 세션을 유지하기 위해 client에 쿠키를 설정.
	client.Jar.SetCookies(ridiUrl, loginResp.Cookies())
	loginResp.Body.Close()
	fmt.Println("쿠키 설정 완료")

	historyUri := "https://ridibooks.com/order/history?page="

	const maxTries int = 100
	totalPrice := 0

	for i := 1; i <= maxTries; i++ {
		historyResp, err := client.Get(historyUri + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err, ": history response error")
			panic(err)
		}

		html, err := goquery.NewDocumentFromResponse(historyResp)
		if err != nil {
			fmt.Println(err, ": goquery parsing error")
			panic(err)
		}

		buyHistory := html.Find("span.museo_sans")
		if len(buyHistory.Nodes) == 0 {
			break
		}

		for _, node := range buyHistory.Nodes {
			price, _ := strconv.Atoi(node.Data)
			totalPrice += price
		}
	}

	fmt.Printf("총 구매 금액 : %d\n", totalPrice)
}
