package tutorial

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func TestWebCrawl_DepthNull(t *testing.T) {
	wg := new(sync.WaitGroup)
	fp := new(F)

	fp.ok = make(map[string]bool)
	Crawl("http://golang.org/", 0, fp, wg, fetcher)
	wg.Wait()

	if fmt.Sprint(fp.e) != "depth cannot be smaller equal zero" {
		t.Error("Return on depth smaller equal zero did not work.")
	}
}

func TestWebCrawl(t *testing.T) {
	wg := new(sync.WaitGroup)
	fp := new(F)

	fp.ok = make(map[string]bool)
	Crawl("http://golang.org/", 4, fp, wg, fetcher)
	wg.Wait()

	m := map[string]bool{
		"http://golang.org/":         true,
		"http://golang.org/cmd/":     true,
		"http://golang.org/pkg/":     true,
		"http://golang.org/pkg/os/":  true,
		"http://golang.org/pkg/fmt/": true,
	}

	if !reflect.DeepEqual(fp.ok, m) {
		t.Error("Webcrawler did not produce expected output.")
	}
}
