package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://areyoukiddingme.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://baidu.com",
		"waat://areyoukiddingme.com",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://baidu.com": true,
		"waat://areyoukiddingme.com": false,
	}

	got := checkWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}