package db

import (
	"context"
	"database/sql"
	"fmt"
	"golang_testing_grpc/pkg/config"

	_ "github.com/lib/pq"
)

type IDatabaseInterface interface {
	DeleteFromTableWithCondition(ctx context.Context, tableName string, condition string, args ...interface{}) error
}

type Database struct {
	db *sql.DB
}

func NewDatabase(uri string) (*Database, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Database connection not tested: %v", err)
	}
	return &Database{db: db}, nil
}

func (d *Database) CloseDatabase() error {
	if d.db != nil {
		err := d.db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) DeleteFromTableWithCondition(ctx context.Context, tableName string, condition string, args ...interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeOut)
	defer cancel()

	query := fmt.Sprintf(`DELETE FROM %s WHERE %s`, tableName, condition)
	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetFromTableWithCondition(ctx context.Context, tableName string, condition string, model interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeOut)
	defer cancel()

	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s`, tableName, condition)
	err := d.db.QueryRowContext(ctx, query).Scan(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (d *Database) GetFromTableWithConditionInList(ctx context.Context, tableName string, condition string, model interface{}) ([]interface{}, error) {
	var newList []interface{}
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeOut)
	defer cancel()

	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s`, tableName, condition)
	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(model); err != nil {
			return nil, err
		}
		newList = append(newList, model)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return newList, nil
}
