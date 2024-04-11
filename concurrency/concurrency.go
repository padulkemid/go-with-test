package concurrency

type WebsiteChecker func(string) bool
type result struct {
  string
  bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
  resChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
      resChan <- result{u, wc(u)}
		}(url)
	}

  for i := 0; i < len(urls); i++ {
    r := <- resChan
    res[r.string] = r.bool
  }

	return res
}
