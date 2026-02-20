package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"http://www.google.com":   true,
		"http://www.facebook.com": true,
		"waat://furhurterwe.geds": false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"waat://furhurterwe.geds",
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
