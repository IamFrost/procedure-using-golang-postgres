package utils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var logrusLog = logrus.New()

// CreateConnection creates database connection
func CreateConnection() *pgxpool.Pool {
	logrusLog.Infof("Connecting....")
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/database3?sslmode=disable")
	if err != nil {
		logrusLog.Errorf("Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}