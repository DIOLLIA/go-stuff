package selector

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare response time of the servers and return the fastest url", func(t *testing.T) {
		/* fixme because this is real servers we can't be sure which will be faster at what time
		slowURL := "http://www.facebook.com"
		fastURL := "http//:go.dev"

		expected := slowURL
		result, _ := Racer(slowURL, fastURL)

		if expected != result {
			t.Errorf("expected %q, got %q", expected, result)
		}*/
		slowServer := createServerWithDelay(20 * time.Microsecond)
		fastServer := createServerWithDelay(0 * time.Microsecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		result, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		if expected != result {
			t.Errorf("expected %q, got %q", expected, result)
		}

	})
	t.Run("returns an error if server not responding within 10 seconds", func(t *testing.T) {
		serverA := createServerWithDelay(21 * time.Millisecond)

		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
