package dialect

import (
	"github.com/yungsem/gotool/pattern"
	"github.com/yungsem/gotool/strs"
)

type MySQL struct {
}

func (m *MySQL) Eq(k string) string {
	return Eq(k)
}

func (m *MySQL) Gt(k string) string {
	return Gt(k)
}

func (m *MySQL) Ge(k string) string {
	return Ge(k)
}

func (m *MySQL) Lt(k string) string {
	return Lt(k)
}

func (m *MySQL) Le(k string) string {
	return Le(k)
}

func (m *MySQL) Like(k string) string {
	return strs.Builder().Append(k).Append(pattern.Space).Append(OpLike).Append(" %?%").String()
}

func (m *MySQL) In(k string) string {
	return In(k)
}