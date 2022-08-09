package fetcher

import (
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
)

func TestFetchJson(t *testing.T) {
	result, err := FetchJson(1)
	if err != nil {
		t.Error(err)
	}

	ioutil.WriteFile("test.json", result, 0666)
}

func TestFetchEvents(t *testing.T) {
	events, err := FetchEvents(10)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(events)
}

func TestConcurrentFetchJson(t *testing.T) {
	goroutineNum, loopNum := 100, 10
	waitGroup := sync.WaitGroup{}

	for i := 0; i < goroutineNum; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()

			for j := 0; j < loopNum; j++ {
				_, err := FetchJson(100)
				if err != nil {
					t.Error(err)
				}
			}
		}()
	}
	waitGroup.Wait()
}
