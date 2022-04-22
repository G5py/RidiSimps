package main

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var logger *log.Logger

func loginData() (id string, password string) {
	fmt.Print("ID : ")
	fmt.Scanln(&id)
	fmt.Print("PASSWORD : ")
	fmt.Scanln(&password)

	return
}

func loginReq(id string, pw string) *http.Request {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	fieldNames := []string{ // 리디북스 로그인 request의 body field들.
		"user_id",
		"password",
		"cmd",
		"return_url",
		"return_query_string",
		"device_id",
		"msg"}

	w.WriteField(fieldNames[0], id)
	w.WriteField(fieldNames[1], pw)
	w.WriteField(fieldNames[2], "login")

	for _, fn := range fieldNames[3:] {
		w.CreateFormField(fn)
	}

	w.Close()
	req, err := http.NewRequest("POST", "https://ridibooks.com/account/action/login", &body)
	if err != nil {
		logger.Panic("loginReq() : failed to http.NewRequest() for the login-POST-request.")
	}
	req.Header.Set("Content-Type", "multipart/form-data; boundary="+w.Boundary())

	return req
}

func getCost(client *http.Client) int {
	uri := getUri()

	const maxTries = 5
	totalCost := 0
	for i := 1; i <= maxTries; i++ {
		resp, err := client.Get(uri())
		if err != nil {
			logger.Panic("gotCost() : failed to Get a buy-history page " + strconv.Itoa(i))
		}

		cost := sumBuyTable(resp)
		if cost == 0 {
			break
		}

		totalCost += cost
	}

	return totalCost
}

func getUri() func() string {
	i := 0
	uri := "https://ridibooks.com/order/history?page="
	return func() string {
		i++
		return uri + strconv.Itoa(i)
	}

}

func sumBuyTable(resp *http.Response) int {
	cost := 0
	sels := parseResp(resp)
	sels.Each(func(i int, s *goquery.Selection) {
		price, err := strconv.Atoi(strings.ReplaceAll(s.Text(), ",", ""))
		if err != nil {
			logger.Panic("sumBuyTable() : failed to Atoi().")
		}
		cost += price
	})

	return cost
}

func parseResp(resp *http.Response) *goquery.Selection {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Panic("parseResp() : error occured from goquery.NewDocumentFromReader().")
	}

	sels := doc.Find("#page_buy_history") // #page_buy_history table tbody tr
	return sels.Find("td.main_value span")
}

func main() {
	logger = log.New(os.Stdout, "Ridisimps : ", log.LstdFlags)
	jar, err := cookiejar.New(nil)
	if err != nil {
		logger.Panic("main() : failed to New cookijar.")
	}
	client := &http.Client{Jar: jar}

	loginResp, err := client.Do(loginReq(loginData()))
	if err != nil {
		logger.Panic("main() : failed to Do login request.")
	}
	defer loginResp.Body.Close()

	fmt.Println("총 결제 금액 : " + strconv.Itoa(getCost(client)))
}
