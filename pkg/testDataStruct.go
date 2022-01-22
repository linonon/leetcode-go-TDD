package pkgu

import "fmt"

type TestDataStruct struct {
	Num       int
	TestData  interface{}
	TestData2 interface{}
	Want      interface{}
	Want2     interface{}
}

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

type TestListNodeStruct struct {
	Num int

	ListNode1 *ListNode
	ListNode2 *ListNode
	ListNode3 *ListNode

	Want *ListNode
}
