// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
