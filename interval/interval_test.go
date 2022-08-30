package interval

import "testing"

func TestLoop(t *testing.T) {
    err := dailySync()
    if err != nil {
        t.Error(err)
    }
    err = languageSync()
    if err != nil {
        t.Error(err)
    }
}
