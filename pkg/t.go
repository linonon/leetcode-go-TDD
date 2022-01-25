package pkgu

import (
	"fmt"
)

// T is a struct for Test.
type T struct {
	Num  int
	Text string

	Val    interface{} // Value
	Want   interface{} // Want
	Target int

	Val1  interface{} // Value1
	Val2  interface{} // Value2
	Want1 interface{} // Want1
	Want2 interface{} // Want2

	// If len(Input) >= 3, Use V and W.
	V []interface{} // Muti Value
	W []interface{} // Muti Want
}

// Name of testing
func (t *T) Name() string {
	if t.Num != 0 {
		return fmt.Sprintf("Test %v", t.Num)
	} else if t.Text != "" {
		return fmt.Sprintf("Test %v", t.Text)
	}

	return "dosen't set num or text"
}

// SetM setting muti interface{} value, return []interface{}
func SetM(a ...interface{}) []interface{} {
	return a
}
