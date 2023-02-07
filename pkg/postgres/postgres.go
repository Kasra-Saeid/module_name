package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	timeOut       time.Duration
	retryAttempts int
	maxPoolSize   int

	SqlBuilder goqu.DialectWrapper
	PgPool     *pgxpool.Pool
}

const (
	_timeOut       = time.Second
	_retryAttempts = 10
	_maxPoolSize   = 10
)

func New(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		timeOut:       _timeOut,
		retryAttempts: _retryAttempts,
		maxPoolSize:   _maxPoolSize,
		SqlBuilder:    goqu.DialectWrapper{},
		PgPool:        &pgxpool.Pool{},
	}

	for _, opt := range opts {
		opt(pg)
	}

	pg.SqlBuilder = goqu.Dialect("postgres")

	pgxConfig, err := pgxpool.ParseConfig(url)

	if err != nil {
		return nil, fmt.Errorf("pkg - postgres - ParseConfig: %w", err)
	}

	pgxConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.retryAttempts > 0 {
		pg.PgPool, err = pgxpool.NewWithConfig(context.Background(), pgxConfig)
		if err == nil {
			break
		}

		log.Println("trying to connect to database, remaining times:", pg.retryAttempts)

		time.Sleep(pg.timeOut)

		pg.retryAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("pkg - postgres - connection attempts == 0: %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.PgPool != nil {
		p.PgPool.Close()
	}
}
