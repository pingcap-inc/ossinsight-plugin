package fetcher

import (
	"sync"
	"testing"
)

func TestFetchJson(t *testing.T) {
	_, err := FetchJson()
	if err != nil {
		t.Error(err)
	}
}

func TestConcurrentFetchJson(t *testing.T) {
	goroutineNum, loopNum := 100, 10
	waitGroup := sync.WaitGroup{}

	for i := 0; i < goroutineNum; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()

			for j := 0; j < loopNum; j++ {
				_, err := FetchJson()
				if err != nil {
					t.Error(err)
				}
			}
		}()
	}
	waitGroup.Wait()
}
