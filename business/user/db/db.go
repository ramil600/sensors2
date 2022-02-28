package db

import (
	"context"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Scheme   string
	Host     string
	User     string
	Password string
	Path     string
}

var DBcfg = DBConfig{
	Scheme:   "postgres",
	Host:     "0.0.0.0:5432",
	User:     "postgres",
	Password: "postgres",
	Path:     "postgres",
}

func Open(cfg DBConfig) (*sqlx.DB, error) {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   cfg.Scheme,
		Host:     cfg.Host,
		User:     url.UserPassword(cfg.User, cfg.Password),
		Path:     cfg.Path,
		RawQuery: q.Encode(),
	}
	db, err := sqlx.Open(cfg.Scheme, u.String())
	if err != nil {
		return nil, err
	}

	return db, nil

}

func (s Store) Ping(ctx context.Context) error {

	const chk = `SELECT true`
	var tmp bool
	err := s.DB.QueryRowContext(ctx, chk).Scan(&tmp)
	if err != nil {
		return err
	}
	return nil

}
