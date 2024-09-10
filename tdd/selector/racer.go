package selector

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecTimeout = 10 * time.Second

/*
no concurrency
func Racer(url1, url2 string) (winner string) {

		url1Duration := measureResponseTime(url1)
		url2Duration := measureResponseTime(url2)

		if url1Duration < url2Duration {
			return url1
		}

		return url2
	}
*/
func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select { //select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout): //very handy function when using select.
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} { //chan struct{} is the smallest data type available from a memory perspective
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
