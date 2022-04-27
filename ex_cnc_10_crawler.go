/* Exercise: Web Crawler
 * See https://go.dev/tour/concurrency/10
 * Description: Implement a parallel web crawler without fetching the same URL twice
 */

package main

import (
    "fmt"
    "sync"
)

// Visited URLs cache
type Visited struct {
    url map[string]bool
    mux sync.Mutex
}

// Set visited flag for an URL to true
func (v *Visited) Set(u string) {
    v.mux.Lock()
    v.url[u] = true
    //fmt.Printf("setting URL '%s' visited.\n", u)
    v.mux.Unlock()
}

// Get visited flag for an URL
func (v *Visited) Get(u string) bool {
    v.mux.Lock()
    defer v.mux.Unlock()
    //fmt.Printf("URL '%s' is visited: %v.\n", u, v.url[u])
    return v.url[u]
}

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    defer wg.Done()
    if depth <= 0 {
        //fmt.Println("depth limit hit. exiting")
        return
    }
    // Don't fetch the same URL twice
    if ( vu.Get(url) == true ) {
        //fmt.Println("already crawled. exiting")
        return
    }
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    vu.Set(url)
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        // Fetch URLs in parallel
        //fmt.Printf("going further: %s\n", u)
        wg.Add(1)
        go Crawl(u, depth-1, fetcher)
    }
    return
}

var vu = Visited{url: make(map[string]bool)}
var wg sync.WaitGroup

func main() {
    wg.Add(1)
    Crawl("https://go.dev/", 4, fetcher)
    wg.Wait()

    //fmt.Println("Done.")
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "https://go.dev/": &fakeResult{
        "The Go Programming Language",
        []string{
            "https://go.dev/pkg/",
            "https://go.dev/cmd/",
        },
    },
    "https://go.dev/pkg/": &fakeResult{
        "Packages",
        []string{
            "https://go.dev/",
            "https://go.dev/cmd/",
            "https://go.dev/pkg/fmt/",
            "https://go.dev/pkg/os/",
        },
    },
    "https://go.dev/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "https://go.dev/",
            "https://go.dev/pkg/",
        },
    },
    "https://go.dev/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "https://go.dev/",
            "https://go.dev/pkg/",
        },
    },
}

