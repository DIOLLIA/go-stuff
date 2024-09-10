package concurrency

import "sync"

type WebsiteCheckerDeclared func(string) bool

/*
	func CheckWebsitesDeclared(wc WebsiteCheckerDeclared, urls []string) map[string]bool {
		results := make(map[string]bool)
		wg := sync.WaitGroup{}
		mu := sync.Mutex{}

		for _, url := range urls {
			wg.Add(1)
			go func(uri string) {
				defer wg.Done()
				result := wc(uri)
				mu.Lock()
				results[uri] = result
				mu.Unlock()
			}(url)
		}
		wg.Wait()
		return results
	}
*/
func CheckWebsitesDeclared(wc WebsiteCheckerDeclared, urls []string) map[string]bool {
	results := make(map[string]bool)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, url := range urls {
		checkWebSite(&wg, &mu, results, wc, url)
	}
	wg.Wait()
	return results
}

func checkWebSite(wg *sync.WaitGroup, mu *sync.Mutex, results map[string]bool, wc WebsiteCheckerDeclared, url string) {
	wg.Add(1)
	defer wg.Done()
	result := wc(url)
	mu.Lock()
	results[url] = result
	mu.Unlock()
}
