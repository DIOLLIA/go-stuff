package concurrency

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		expectedCount := 3
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, expectedCount, counter)
	})
	t.Run("run concurrent", func(t *testing.T) {
		expectedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, expectedCount, counter)
	})
}

func assertCounter(t testing.TB, expectedCounter int, result *Counter) {
	t.Helper()
	if expectedCounter != result.Value() {
		t.Errorf("expected %d, result %d", expectedCounter, result.Value())
	}
}
