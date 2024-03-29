package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 33306
	// GitHub Actions always set CI to true
	// https://docs.github.com/en/actions/learn-github-actions/variables#default-environment-variables
	if _, defined := os.LookupEnv("CI"); defined {
		port = 3306
	}
	c := mysql.Config{
		User:                 "todo",
		Passwd:               "todo",
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("127.0.0.1:%d", port),
		DBName:               "todo",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sql.Open(
		"mysql",
		c.FormatDSN(),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "mysql")
}

func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	host := "127.0.0.1"
	port := 36379
	if _, defined := os.LookupEnv("CI"); defined {
		port = 6379
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect redis: %s", err)
	}
	return client
}
