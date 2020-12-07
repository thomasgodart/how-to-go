package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSimple(t *testing.T) {

	tests := []struct {
		url  string
		file string
		str  string
	}{
		{"", "test/serve-.txt", ``},
		{"aaa", "test/serve-aaa.txt", ``},
		{"bbb", "test/serve-bbb.txt", ``},
	}

	for _, test := range tests {

		// load expected output from files, to avoid editor formatting
		if binary, err := ioutil.ReadFile(test.file); err == nil {
			test.str = string(binary)
		}

		simple := Simple(test.url)

		if len(simple) != len(test.str) {
			t.Error(fmt.Sprintf(`len(simple) != len(test.str): %d != %d`, len(simple), len(test.str)))
		}
		if simple != test.str {
			t.Error(cmp.Diff(simple, test.str))
		}
	}
}

func Simple(url string) string {

	res, err := http.Get("http://how-srv-simple:8080/" + url)
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
