package risingwave

import (
	"strconv"
	"testing"
	"time"
)

func TestInsertMessage(t *testing.T) {
	id := "id" + strconv.Itoa(time.Now().Nanosecond())
	alreadyExist, err := InsertMessage(id, 0, 0, time.Now(), "programLanguage")
	if err != nil {
		t.Error(err)
	}

	if alreadyExist {
		t.Error("alreadyExist should be false")
	}

	alreadyExist, err = InsertMessage(id, 0, 0, time.Now(), "programLanguage")
	if err != nil {
		t.Error(err)
	}

	if !alreadyExist {
		t.Error("alreadyExist should be true")
	}
}
