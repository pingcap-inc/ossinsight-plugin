package redis

import (
	"strconv"
	"testing"
	"time"
)

func TestExistsAndSet(t *testing.T) {
	key := strconv.Itoa(time.Now().Nanosecond())
	exist, err := ExistsAndSet(key)
	if err != nil {
		t.Error(err)
		return
	}

	if exist {
		t.Errorf("%s should not exists", key)
		return
	}

	exist, err = ExistsAndSet(key)
	if err != nil {
		t.Error(err)
		return
	}

	if !exist {
		t.Errorf("%s should exists", key)
		return
	}
}
