package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestParseProfile(t *testing.T) {
	body, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	results := ParseProfile(body)
	fmt.Println(results)
}
