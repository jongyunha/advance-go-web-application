package entity

import "github.com/jmoiron/sqlx"

type TransactionManager interface {
	Do(func(tx *sqlx.Tx) error) error
}

type SqlxTransactionManager struct {
	db *sqlx.DB
}

func NewSqlxTransactionManager(db *sqlx.DB) *SqlxTransactionManager {
	return &SqlxTransactionManager{db: db}
}

func (s *SqlxTransactionManager) Do(fn func(tx *sqlx.Tx) error) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	return tx.Commit()
}
