package interval

import "testing"

func TestLoop(t *testing.T) {
	err := todayCountSync()
	if err != nil {
		t.Error(err)
	}

	err = dailySync()
	if err != nil {
		t.Error(err)
	}

	err = yearCountSync()
	if err != nil {
		t.Error(err)
	}
}
