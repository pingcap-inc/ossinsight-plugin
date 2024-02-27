package risingwave

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pingcap-inc/ossinsight-plugin/tidb"
	"time"
)

func UpsertEventDaily(events []tidb.DailyEvent) error {
	initRisingWave()

	batch := pgx.Batch{}
	for _, event := range events {
		batch.Queue(`
			INSERT INTO t_event_daily 
			    (day, pr, open, merge, close, dev) 
			VALUES ($1, $2, $3, $4, $5)`,
			event.EventDay, event.OpenedPRs+event.ClosedPRs+event.MergedPRs,
			event.OpenedPRs, event.MergedPRs, event.ClosedPRs, event.Developers)
	}

	_, err := risingWave.SendBatch(context.Background(), &batch).Exec()
	return err
}

func IncreaseEventDaily(pr, open, merge, close, dev int) error {
	initRisingWave()

	_, err := risingWave.Exec(context.Background(), `
		UPDATE t_event_daily
		SET pr = pr + $1, open = open + $2, merge = merge + $3, close = close + $4, dev = dev + $5
		WHERE day = $6`,
		time.Now().Format("2006-01-02"), pr, open, merge, close, dev)

	return err
}
