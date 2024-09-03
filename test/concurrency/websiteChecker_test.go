package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "http://www.google.com"
}
func slowStubWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://blog.gypsydave5.com",
		"http://www.bing.com",
	}
	expected := map[string]bool{
		"http://www.google.com":      false,
		"http://blog.gypsydave5.com": true,
		"http://www.bing.com":        true,
	}

	result := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

/*
No concurrency (commented block of code)
BenchmarkCheckWebsites-12              1        3095533000 ns/op
PASS
ok      workout/concurrency     4.328s

with concurrency
BenchmarkCheckWebsites-12             46          30999000 ns/op
PASS
ok      workout/concurrency     2.584s
*/
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
