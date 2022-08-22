package tidb

import (
	"database/sql"
	"fmt"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	initDBOnce sync.Once
	tidb       *sql.DB
)

func createDB() {
	tidbConfig := config.GetReadonlyConfig().Tidb
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		tidbConfig.User, tidbConfig.Password, tidbConfig.Host, tidbConfig.Port, tidbConfig.Db)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	tidb = db
}
