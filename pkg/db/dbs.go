package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type IDatabaseInterface interface {
	JustQueryForList(ctx context.Context, tableName string, query string) (*sql.Rows, error)
	InsertInto(tableName string, dataPlace string, insertData string, args ...interface{}) (sql.Result, error)
	UpdateData(tableName string, newData string, whereData string, values ...interface{}) (sql.Result, error)
	DeleteData(tableName string, whereData string) (sql.Result, error)
	QueryRow(tableName string, query string) *sql.Row
	GetLimit(pageLimit int) string
	WhereData(query string) string
	GetOffset(pageOffset int) string
}

type Database struct {
	db *sql.DB
}

const (
	UsersTable     = "user"
	ProductTable   = "product"
	OrderTable     = "order"
	OrderLineTable = "order_line"
)

const (
	Id    = "id"
	Order = "order"
)

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

func (d *Database) JustQueryForList(ctx context.Context, tableName string, query string) (*sql.Rows, error) {
	q := fmt.Sprintf(`SELECT * FROM %s %s`, tableName, query)
	rows, err := d.db.Query(q)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) QueryRow(tableName string, query string) *sql.Row {
	q := fmt.Sprintf(`SELECT * FROM %s %s`, tableName, query)
	row := d.db.QueryRow(q)
	return row
}

func (d *Database) WhereData(query string) string {
	q := fmt.Sprintf(" WHERE %s", query)
	return q
}

func (d *Database) GetLimit(pageLimit int) string {
	q := fmt.Sprintf(" LIMIT %d", pageLimit)
	return q
}

func (d *Database) GetOffset(pageOffset int) string {
	q := fmt.Sprintf(" OFFSET %d", pageOffset)
	return q
}

func (d *Database) InsertInto(tableName string, dataPlace string, insertData string, args ...interface{}) (sql.Result, error) {
	q := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, tableName, dataPlace, insertData)
	rows, err := d.db.Exec(q, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (d *Database) UpdateData(tableName string, newData string, whereData string, values ...interface{}) (sql.Result, error) {
	q := fmt.Sprintf(`UPDATE %s SET %s WHERE %s`, tableName, newData, whereData)
	rows, err := d.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) DeleteData(tableName string, whereData string) (sql.Result, error) {
	q := fmt.Sprintf(`DELETE FROM %s WHERE %s`, tableName, whereData)
	row, err := d.db.Exec(q)
	if err != nil {
		return nil, err
	}
	return row, nil
}
