package ridisimps

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParsePurchaseHistories(responses []*http.Response) [][]int {
	costs := make([][]int, 0, 20)

	for _, response := range responses {
		costs = append(costs, ParsePurchaseHistoryToCosts(response))
	}

	return costs
}

func ParsePurchaseHistoryToCosts(response *http.Response) []int {
	doc, _ := goquery.NewDocumentFromReader(response.Body)

	sels := doc.Find("#page_buy_history").Find(".museo_sans")

	costs := make([]int, 0, 15)
	sels.Each(func(i int, s *goquery.Selection) {
		cost, _ := strconv.Atoi(strings.ReplaceAll(s.Text(), ",", ""))
		costs = append(costs, cost)
	})

	return costs
}
