package schema

import "github.com/yungsem/dbx/dialect"

type Condition struct {
	Dialect dialect.Dialect
	Values  []interface{}
	Ops     map[string]struct{}
	expr    []string
}

func (c *Condition) Eq(k string, v interface{}) *Condition {
	return c.build("=", v, c.Dialect.Eq(k))
}

func (c *Condition) Gt(k string, v interface{}) *Condition {
	return c.build(">", v, c.Dialect.Gt(k))
}

func (c *Condition) Ge(k string, v interface{}) *Condition {
	return c.build(">=", v, c.Dialect.Ge(k))
}

func (c *Condition) Lt(k string, v interface{}) *Condition {
	return c.build("<", v, c.Dialect.Lt(k))
}

func (c *Condition) Le(k string, v interface{}) *Condition {
	return c.build("<=", v, c.Dialect.Le(k))
}

func (c *Condition) In(k string, v interface{}) *Condition {
	return c.build("IN", v, c.Dialect.In(k))
}

func (c *Condition) build(op string, v interface{}, expr string) *Condition {
	c.Values = append(c.Values, v)
	c.Ops[op] = struct{}{}
	c.expr = append(c.expr, expr)
	return c
}
