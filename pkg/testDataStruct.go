package pkgu

type TestDataStruct struct {
	Num       int
	TestData  interface{}
	TestData2 interface{}
	Want      interface{}
	Want2     interface{}
}

type TestListNodeStruct struct {
	Num int

	ListNode1 *ListNode
	ListNode2 *ListNode
	ListNode3 *ListNode

	Want *ListNode
}
