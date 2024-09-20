package main

import (
	"RidiSimps/ridisimps"
	"fmt"
	"strconv"
)

func main() {
	client := ridisimps.Login()
	response := ridisimps.RequestPurchaseHistories(client)
	costs := ridisimps.ParsePurchaseHistories(response)
	totalCost := ridisimps.SumIntSlices(costs)
	fmt.Println("총 결제 금액 : " + ridisimps.PutCommasAtNumber(strconv.Itoa(totalCost)) + "원")
}
