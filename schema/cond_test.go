package schema

import (
	"fmt"
	"testing"
)

func TestCond_Eq(t *testing.T) {
	var cond *Condition
	cond.Eq("name", "yangsen").Eq("age", 32)
	fmt.Printf("%v\n", cond)
}
