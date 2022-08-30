package tidb

import (
    "github.com/pingcap-inc/ossinsight-plugin/config"

    _ "github.com/go-sql-driver/mysql"
)

type DailyEvent struct {
    EventDay string `json:"eventDay"`
    Events   int    `json:"events"`
}

func QueryThisYearDailyEvent() ([]DailyEvent, error) {
    initDBOnce.Do(createDB)

    tidbConfig := config.GetReadonlyConfig().Tidb
    rows, err := tidb.Query(tidbConfig.Sql.EventsDaily)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []DailyEvent
    for rows.Next() {
        event := DailyEvent{}
        err = rows.Scan(&event.EventDay, &event.Events)
        if err == nil {
            result = append(result, event)
        } else {
            return result, err
        }
    }

    return result, nil
}

// QueryThisYearPRDailyEvent return opened event, merged event, and error
func QueryThisYearPRDailyEvent() ([]DailyEvent, []DailyEvent, error) {
    initDBOnce.Do(createDB)

    tidbConfig := config.GetReadonlyConfig().Tidb
    rows, err := tidb.Query(tidbConfig.Sql.PrDaily)
    if err != nil {
        return nil, nil, err
    }
    defer rows.Close()

    var openEvent []DailyEvent
    var mergeEvent []DailyEvent
    for rows.Next() {
        event := DailyEvent{}
        action := ""
        err = rows.Scan(&action, &event.EventDay, &event.Events)
        if err == nil {
            if action == "closed" {
                mergeEvent = append(mergeEvent, event)
            } else {
                openEvent = append(openEvent, event)
            }
        } else {
            return openEvent, mergeEvent, err
        }
    }

    return openEvent, mergeEvent, nil
}

func QueryThisYearDeveloperDailyEvent() ([]DailyEvent, error) {
    initDBOnce.Do(createDB)

    tidbConfig := config.GetReadonlyConfig().Tidb
    rows, err := tidb.Query(tidbConfig.Sql.PrDeveloperDaily)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []DailyEvent
    for rows.Next() {
        event := DailyEvent{}
        err = rows.Scan(&event.EventDay, &event.Events)
        if err == nil {
            result = append(result, event)
        } else {
            return result, err
        }
    }

    return result, nil
}
