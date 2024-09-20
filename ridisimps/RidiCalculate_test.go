package ridisimps

import (
	"testing"
)

func TestPutCommasAtNumber(t *testing.T) {
	testCase := make(map[string]string, 0)
	testCase["1000"] = "1,000"
	testCase["10000"] = "10,000"
	testCase["100000"] = "100,000"

	for number, expected := range testCase {
		actual := PutCommasAtNumber(number)
		if expected != actual {
			t.Error("Expected is " + expected + " " + "but actual is " + actual)
		}
	}
}
