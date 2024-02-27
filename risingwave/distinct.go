package risingwave

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

func EventIDExists(id string) (bool, error) {
	initRisingWave()

	ctx := context.Background()

	existCount := 0
	err := risingWave.QueryRow(ctx, `
		SELECT COUNT(*) FROM t_github_prs
		WHERE id = @id
	`, pgx.NamedArgs{"id": id}).Scan(&existCount)

	if err != nil {
		return false, err
	}

	return existCount != 0, nil
}

func DevIDInsertAndTestExists(developer int64) (bool, error) {
	initRisingWave()

	ctx := context.Background()

	existCount := 0
	err := risingWave.QueryRow(ctx, `
		SELECT COUNT(*) FROM t_developer_id
		WHERE developer = @developer
		AND create_time >= date_trunc('year', CURRENT_TIMESTAMP)
	`, pgx.NamedArgs{"developer": developer}).Scan(&existCount)

	if err != nil {
		return false, err
	}

	createTime := time.Unix(time.Now().Unix(), 0)
	_, err = risingWave.Exec(ctx, `
		INSERT INTO t_developer_id 
		VALUES (@developer, @create_time)
	`, pgx.NamedArgs{
		"developer":   developer,
		"create_time": createTime,
	})

	if err != nil {
		return false, err
	}

	risingWave.Exec(ctx, "FLUSH")

	return existCount != 0, nil
}

func RepoIDInsertAndTestExists(repository int64) (bool, error) {
	initRisingWave()

	ctx := context.Background()

	existCount := 0
	err := risingWave.QueryRow(ctx, `
		SELECT COUNT(*) FROM t_repository_id
		WHERE repository = @repository
		AND create_time >= date_trunc('year', CURRENT_TIMESTAMP)
	`, pgx.NamedArgs{"repository": repository}).Scan(&existCount)

	if err != nil {
		return false, err
	}

	createTime := time.Unix(time.Now().Unix(), 0)
	_, err = risingWave.Exec(ctx, `
		INSERT INTO t_repository_id 
		VALUES (@repository, @create_time)
	`, pgx.NamedArgs{
		"repository":  repository,
		"create_time": createTime,
	})

	if err != nil {
		return false, err
	}

	risingWave.Exec(ctx, "FLUSH")

	return existCount != 0, nil
}
