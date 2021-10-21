package dialect

import (
	"github.com/yungsem/gotool/pattern"
	"github.com/yungsem/gotool/strs"
)

const (
	DriverMySQL     = "mysql"
	DriverSqlServer = "sqlserver"
	DriverOracle    = "oracle"
)

type Dialect interface {
	operator
}

const (
	OpEq   = "="
	OpIn   = "IN"
	OpGt   = ">"
	OpGe   = ">="
	OpLt   = "<"
	OpLe   = "<="
	OpLike = "LIKE"
)

type operator interface {
	Eq(k string) string
	Gt(k string) string
	Ge(k string) string
	Lt(k string) string
	Le(k string) string
	Like(k string) string
	In(k string) string
}

func Eq(k string) string {
	return strs.Builder().Append(k).Append(pattern.Space).Append(OpEq).Append("?").String()
}

func Gt(k string) string {
	return strs.Builder().Append(k).Append(pattern.Space).Append(OpGt).Append("?").String()
}

func Ge(k string) string {
	return strs.Builder().Append(k).Append(pattern.Space).Append(OpGe).Append("?").String()
}

func Lt(k string) string {
	return strs.Builder().Append(k).Append(pattern.Space).Append(OpLt).Append("?").String()
}

func Le(k string) string {
	return strs.Builder().Append(k).Append(pattern.Space).Append(OpLe).Append("?").String()
}

func In(k string) string {
	return strs.Builder().Append(k).Append(pattern.Space).Append(OpIn).Append(" (?)").String()
}
