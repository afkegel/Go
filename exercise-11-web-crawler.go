package tutorial

import (
	"errors"
	"sync"
)

// Fetcher abstracts the fetch method provided to Crawl.
type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// F is a lockable map that has to be provided to Crawl to cache fetched urls.
type F struct {
	ok    map[string]bool
	e     error
	mutex sync.RWMutex
}

// fetched and setFetched are the getter and setter methods on F
func (f *F) fetched(url string) (bool, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	if f.e != nil {
		return false, f.e
	}

	return f.ok[url], nil
}

func (f *F) setFetched(url string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.ok[url] = true
}

// Crawl recursively crawls url to a maximum of depth. The caller has to
// provide a lockable map *F, that keeps track of scanned urls, a wait group
// *sync.Waitgroup to count goroutines, and an implementation of the Fetcher
// interface. Crawl fetches urls concurrently and skips previously fetched urls
// in case of duplicates. The caller is responsible for waiting until the
// goroutines return.
func Crawl(url string, depth int, f *F, wg *sync.WaitGroup, fetcher Fetcher) {
	if depth <= 0 {
		f.e = errors.New("depth cannot be smaller equal zero")
	}

	ok, err := f.fetched(url)
	if err != nil {
		return
	}

	if !ok {
		f.setFetched(url)

		_, urls, err := fetcher.Fetch(url)
		if err != nil {
			return
		}

		for _, u := range urls {
			wg.Add(1)
			go Crawl(u, depth-1, f, wg, fetcher)
			// wg.Done() here and already the first goroutine may release
			// wg.Wait and cause a race
		}
		// wg.Done() here and the wg counter will never go down to zero because
		// some goroutines skip the if clause
	}
	wg.Done()
	return
}
