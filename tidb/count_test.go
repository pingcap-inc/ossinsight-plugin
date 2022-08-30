package tidb

import (
    "fmt"
    "testing"
)

func TestQueryThisYearDeveloperCount(t *testing.T) {
    result, err := QueryThisYearDeveloperCount()
    if err != nil {
        t.Error(err)
    }

    fmt.Println(result)
}
