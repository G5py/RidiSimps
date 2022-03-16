package main

import "fmt"

func loginData() (id string, password string) {
	fmt.Print("ID : ")
	fmt.Scanln(&id)
	fmt.Print("PASSWORD : ")
	fmt.Scanln(&password)

	return id, password
}

func login(id string, password string) {

}

func main() {
	var id string
	var password string

	id, password = loginData()
	login(id, password)

}
