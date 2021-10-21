package dbx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yungsem/dbx/dialect"
	"github.com/yungsem/dbx/schema"
	"github.com/yungsem/gotool/errs"
	"reflect"
)

// DB 是对 *sqlx.DB 的包装
type DB struct {
	*sqlx.DB
	dialect dialect.Dialect
}

// Connect 打开一个 DB 并建立连接，并根据 sql 驱动的类型设置 dialect
func Connect(driverName, dataSourceName string) (*DB, error) {
	// 打开 DB 并建立连接
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	// 根据 sql 驱动的类型设置 dialect
	var dia dialect.Dialect
	if driverName == dialect.DriverMySQL {
		dia = new(dialect.MySQL)
	}

	return &DB{db, dia}, nil
}

func (d *DB) Cond() *schema.Condition {
	var c schema.Condition
	c.Dialect = d.dialect
	c.Ops = make(map[string]struct{})
	return &c
}

func (d *DB) Insert(i interface{}) error {
	sch, err := schema.Parse(i)
	if err != nil {
		return err
	}

	sql := schema.InsertSQL(sch)

	_, err = d.Exec(sql, sch.Values...)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Update(i interface{}) error {
	sch, err := schema.Parse(i)
	if err != nil {
		return err
	}

	sql := schema.UpdateByIdSQL(sch)

	sch.Values = append(sch.Values, sch.Values[0])

	_, err = d.Exec(sql, sch.Values...)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) One(cond *schema.Condition, dest interface{}) error {
	sch, err := schema.ParseColumns(dest)
	if err != nil {
		return err
	}

	sql := schema.QueryByCondSQL(sch, cond)

	err = d.Get(dest, sql, cond.Values...)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) List(cond *schema.Condition, dest interface{}) error {
	t := reflect.TypeOf(dest)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Slice {
		return errs.New(fmt.Sprintf("unsupported dest type %s, only support slice type", t.Kind()))
	}
	t = t.Elem()

	sch, err := schema.ParseColumnsByType(t)
	if err != nil {
		return err
	}

	sql := schema.QueryByCondSQL(sch, cond)

	if _, ok := cond.Ops[dialect.OpIn]; ok {
		s, args, err := sqlx.In(sql, cond.Values...)
		if err != nil {
			return err
		}

		err = d.Select(dest, s, args...)
		if err != nil {
			return err
		}
		return nil
	}

	err = d.Select(dest, sql, cond.Values...)
	if err != nil {
		return err
	}
	return nil
}
