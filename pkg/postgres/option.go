package postgres

import "time"

type Option func(pg *Postgres)

func SetRetryAttempts(retries int) Option {
	return func(pg *Postgres) {
		pg.retryAttempts = retries
	}
}

func SetPostgresTimeOut(timeOut time.Duration) Option {
	return func(pg *Postgres) {
		pg.timeOut = timeOut
	}
}

func SetPostgresPoolSize(poolSize int) Option {
	return func(pg *Postgres) {
		pg.maxPoolSize = poolSize
	}
}
