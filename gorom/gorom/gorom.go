package gorom

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Model interface {
	TableName() string
}

type Repository[T Model] struct {
	DB *sql.DB
}

// NewRepository creates a new repository for the given model type
func NewRepository[T Model](db *sql.DB) *Repository[T] {
	return &Repository[T]{
		DB: db,
	}
}

// All fetches all rows for a given model
func (r *Repository[T]) All() ([]T, error) {
	var temp T
	table := temp.TableName()
	query := fmt.Sprintf("SELECT * FROM %s", table)

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		var entity T
		err := scanRow(&entity, rows)
		if err != nil {
			return nil, err
		}
		results = append(results, entity)
	}

	return results, nil
}

func scanRow[T any](dest *T, rows *sql.Rows) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	val := reflect.ValueOf(dest).Elem()
	if val.Kind() != reflect.Struct {
		return errors.New("destination is not a struct")
	}

	values := make([]interface{}, len(columns))

	for i, col := range columns {
		field := val.FieldByNameFunc(func(name string) bool {
			return strings.EqualFold(name, col)
		})
		if !field.IsValid() || !field.CanSet() {
			var dummy interface{}
			values[i] = &dummy
		} else {
			values[i] = field.Addr().Interface()
		}
	}

	return rows.Scan(values...)
}
