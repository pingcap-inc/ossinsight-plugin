package fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

func FetchJson(size int) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/events?per_page=%d", size))
	if err != nil {
		logger.Error("get events http request error", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("get events http request response code not 200", zap.Int("code", resp.StatusCode))
		return nil, err
	}

	return body, err
}

func FetchEvents(size int) ([]Event, error) {
	payload, err := FetchJson(size)
	if err != nil {
		logger.Error("fetch json error", zap.Error(err))
		return nil, err
	}

	var events []Event
	err = json.Unmarshal(payload, &events)
	if err != nil {
		logger.Error("events unmarshal error", zap.Error(err))
		return nil, err
	}

	return events, nil
}
