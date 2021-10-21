package schema

import (
	"github.com/yungsem/gotool/strs"
	"strings"
)

func InsertSQL(sch *Schema) string {
	sb := strs.Builder().Append("INSERT INTO ").
		Append(sch.Table).Append(" ").
		Append("(").Append(strings.Join(sch.Columns, ",")).Append(")").
		Append(" VALUES ").
		Append("(")

	for i := 0; i < len(sch.Columns); i++ {
		sb.Append("?")
		if i != len(sch.Columns)-1 {
			sb.Append(",")
		}
	}
	sb.Append(")")
	return sb.String()
}

func UpdateByIdSQL(sch *Schema) string {
	sb := strs.Builder().Append("UPDATE ").Append(sch.Table).Append(" SET ")
	for i := 0; i < len(sch.Columns); i++ {
		sb.Append(sch.Columns[i]).Append("=?")
		if i != len(sch.Columns)-1 {
			sb.Append(", ")
		}
	}
	sb.Append(" WHERE id=?")
	return sb.String()
}

func QueryByCondSQL(sch *Schema, cond *Condition) string {
	return strs.Builder().
		Append("SELECT ").Append(strings.Join(sch.Columns, ",")).
		Append(" FROM ").Append(sch.Table).
		Append(" WHERE ").Append(strings.Join(cond.expr, " AND ")).String()
}
