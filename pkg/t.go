package pkgu

import "fmt"

type T struct {
	Num  int
	Text string

	Value interface{}
	Want  interface{}

	Value1 interface{}
	Value2 interface{}
	Value3 interface{}
	Want1  interface{}
	Want2  interface{}
	Want3  interface{}
}

// Testing name
func (t *T) GetNumName() string {
	return fmt.Sprintf("Test %v", t.Num)
}

func (t *T) GetTextName() string {
	return fmt.Sprintf("Test %v", t.Text)
}
