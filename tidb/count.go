package tidb

import (
    "fmt"
    "github.com/pingcap-inc/ossinsight-plugin/config"
)

func QueryThisYearDeveloperCount() (int, error) {
    tidbConfig := config.GetReadonlyConfig().Tidb
    return QueryDeveloperCount(tidbConfig.Sql.PrDeveloperThisYear)
}

func QueryDeveloperCount(sql string) (int, error) {
    initDBOnce.Do(createDB)

    rows, err := tidb.Query(sql)
    if err != nil {
        return 0, err
    }
    defer rows.Close()

    for rows.Next() {
        number := 0
        err = rows.Scan(&number)
        if err == nil {
            return number, nil
        } else {
            return 0, err
        }
    }

    return 0, fmt.Errorf("empty result")
}
