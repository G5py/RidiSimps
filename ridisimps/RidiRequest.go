package ridisimps

import (
	"net/http"
	"strconv"
)

func RequestPurchaseHistories(client *http.Client) []*http.Response {
	const maxTries = 20
	nexturi := makeUriGenerator()
	responseCollection := make([]*http.Response, 0, 20)

	for i := 0; i < maxTries; i++ {
		response, _ := client.Get(nexturi())
		responseCollection = append(responseCollection, response)
	}

	return responseCollection
}

func RequestPurchaseHistoriesAndParse(client *http.Client) []*http.Response {
	const maxTries = 20
	nexturi := makeUriGenerator()
	resonseCollection := make([]*http.Response, 0, 20)

	for i := 0; i < maxTries; i++ {
		response, _ := client.Get(nexturi())

		costs := ParsePurchaseHistoryToCosts(response)
		if IsSliceEmpty(costs) {
			break
		}

		resonseCollection = append(resonseCollection, response)
	}

	return resonseCollection
}

func makeUriGenerator() func() string {
	counter := makeCounter()

	return func() string {
		return "https://ridibooks.com/order/history?page=" + strconv.Itoa(counter())
	}
}

// Count starts at 1.
func makeCounter() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
