package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

type pgStore struct {
	db *sqlx.DB
}

const (
	dbDriver = "postgres"
)

func NewPgStore(db *sqlx.DB) Storer {
	return &pgStore{db}
}

func Init() (s Storer, err error) {
	uri := "postgresql://postgres:1234@localhost:5432/wallet?sslmode=disable"

	conn, err := sqlx.Connect(dbDriver, uri)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot initialize database")
		return
	}

	logger.WithField("uri", uri).Info("Connected to pg database")
	store := NewPgStore(conn)
	return store, nil
}
