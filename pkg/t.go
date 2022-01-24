package pkgu

import (
	"fmt"
)

type T struct {
	Num  int
	Text string

	Val    interface{} // Value
	Want   interface{} // Want
	Target int

	Val1  interface{}
	Val2  interface{}
	Want1 interface{}
	Want2 interface{}

	V []interface{} // Muti Value
	W []interface{} // Muti Want
}

// Testing name
func (t *T) TestName() string {
	if t.Num != 0 {
		return fmt.Sprintf("Test %v", t.Num)
	} else if t.Text != "" {
		return fmt.Sprintf("Test %v", t.Text)
	}

	return "dosen't set num or text"
}

// SetM setting muti interface{} value, return []interface{}
func SetM(a ...interface{}) []interface{} {
	return []interface{}{a}
}
