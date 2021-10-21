package schema

import (
	"github.com/yungsem/gotool/reflects"
	"github.com/yungsem/gotool/strs"
	"go/ast"
	"reflect"
)

type Schema struct {
	Table   string
	Columns []string
	Values  []interface{}
}

func Parse(ent interface{}) (*Schema, error) {
	t := reflect.TypeOf(ent).Elem()
	sch := &Schema{
		Table: strs.ToSnakeCaseLower(t.Name()),
	}

	fields, err := reflects.Fields(ent)
	if err != nil {
		return nil, err
	}

	values, err := reflects.Values(ent)
	if err != nil {
		return nil, err
	}
	sch.Values = values

	for _, f := range fields {
		if !ast.IsExported(f.Name) {
			continue
		}
		sch.Columns = append(sch.Columns, strs.ToSnakeCaseUpper(f.Name))
	}

	return sch, nil
}

func ParseColumns(ent interface{}) (*Schema, error) {
	t := reflect.TypeOf(ent).Elem()
	sch := &Schema{
		Table: strs.ToSnakeCaseLower(t.Name()),
	}

	fields, err := reflects.Fields(ent)
	if err != nil {
		return nil, err
	}

	for _, f := range fields {
		if !ast.IsExported(f.Name) {
			continue
		}
		sch.Columns = append(sch.Columns, strs.ToSnakeCaseUpper(f.Name))
	}

	return sch, nil
}


func ParseColumnsByType(t reflect.Type) (*Schema, error) {
	sch := &Schema{
		Table: strs.ToSnakeCaseLower(t.Name()),
	}

	fields, err := reflects.Fields(reflect.New(t).Elem().Interface())
	if err != nil {
		return nil, err
	}

	for _, f := range fields {
		nn := f.Name
		_ = nn
		if !ast.IsExported(f.Name) {
			continue
		}
		sch.Columns = append(sch.Columns, strs.ToSnakeCaseUpper(f.Name))
	}

	return sch, nil
}