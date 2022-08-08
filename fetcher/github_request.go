package fetcher

import (
    "io/ioutil"
    "net/http"

    "github.com/pingcap/log"
    "go.uber.org/zap"
)

func FetchJson() (string, error) {
    resp, err := http.Get("https://api.github.com/events?per_page=100")
    if err != nil {
        log.Error("get events http request error", zap.Error(err))
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    log.Debug("get events result",
        zap.Int("code", resp.StatusCode),
        zap.ByteString("body", body))
    if err != nil {
        log.Error("get events http request response code not 200", zap.Int("code", resp.StatusCode))
        return "", err
    }

    return string(body), err
}
