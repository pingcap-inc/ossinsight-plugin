package tidb

import (
	"fmt"
	"github.com/pingcap-inc/ossinsight-plugin/config"
)

type PRCount struct {
	Action   string `json:"action"`
	EventNum int    `json:"event_num"`
}

func QueryThisYearDeveloperCount() (int64, error) {
	tidbConfig := config.GetReadonlyConfig().Tidb
	return QueryDeveloperCount(tidbConfig.Sql.PrDeveloperThisYear)
}

func QueryTodayDeveloperCount() (int64, error) {
	tidbConfig := config.GetReadonlyConfig().Tidb
	return QueryDeveloperCount(tidbConfig.Sql.PrDeveloperToday)
}

func QueryDeveloperCount(sql string) (int64, error) {
	initDBOnce.Do(createDB)

	rows, err := tidb.Query(sql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		number := int64(0)
		err = rows.Scan(&number)
		if err == nil {
			return number, nil
		} else {
			return 0, err
		}
	}

	return 0, fmt.Errorf("empty result")
}

func QueryThisYearPRCount() (map[string]int, error) {
	tidbConfig := config.GetReadonlyConfig().Tidb
	return QueryPRCount(tidbConfig.Sql.PrThisYear)
}

func QueryTodayPRCount() (map[string]int, error) {
	tidbConfig := config.GetReadonlyConfig().Tidb
	return QueryPRCount(tidbConfig.Sql.PrToday)
}

func QueryPRCount(sql string) (map[string]int, error) {
	initDBOnce.Do(createDB)

	rows, err := tidb.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		event := PRCount{}
		err = rows.Scan(&event.Action, &event.EventNum)
		if err == nil {
			result[event.Action] = event.EventNum
		} else {
			return result, err
		}
	}

	return result, nil
}
