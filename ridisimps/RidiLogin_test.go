package ridisimps

import (
	"strconv"
	"testing"
)

func TestToLoginRequest(t *testing.T) {

	// Given
	client := makeClient()
	request := LoginData{Id: "asdf", Pw: "asdf"}.toLoginRequest()

	// When
	response, err := client.Do(request)

	// Then
	if err != nil {
		t.Error("An error occurs during request.")
	}

	if response.StatusCode != 200 {
		t.Error("Status code is not 200. Status code is " + strconv.FormatInt(int64(response.StatusCode), 10))
	}

}
