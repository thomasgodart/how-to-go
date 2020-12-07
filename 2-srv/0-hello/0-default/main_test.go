package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSayHello(t *testing.T) {

	tests := []struct {
		url string
		str string
	}{
		{"", "Hello \n"},
		{"World", "Hello World\n"},
	}

	for _, test := range tests {

		say := Say(test.url)

		if say != test.str {
			t.Error(fmt.Sprintf(`say != "%s"`, test.str))
		}
	}
}

func Say(url string) string {

	res, err := http.Get("http://how-srv-hello:8080/" + url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%s", body)

	return s
}
