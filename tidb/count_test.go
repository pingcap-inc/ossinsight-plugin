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

func TestQueryTodayDeveloperCount(t *testing.T) {
	result, err := QueryTodayDeveloperCount()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}

func TestQueryThisYearPRCount(t *testing.T) {
	result, err := QueryThisYearPRCount()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}

func TestQueryTodayPRCount(t *testing.T) {
	result, err := QueryTodayPRCount()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}
