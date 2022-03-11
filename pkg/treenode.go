package pkgu

// NTreeNode is a n cross tree node.
type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

// ToTreeNode transforms []int to *NTreeNode.
func ToTreeNode(s []int) *NTreeNode {
	if len(s) <= 3 {
		return nil
	}

	var result = new(NTreeNode)

	EachChilden := make([][]*NTreeNode, 0)
	mapNode := make(map[int]*NTreeNode)
	for i := 0; i < len(s); i++ {
		Layer := make([]*NTreeNode, 0)
		for i < len(s) && s[i] >= 0 {
			node := &NTreeNode{Val: s[i]}
			Layer = append(Layer, node)
			mapNode[s[i]] = node
			i++
		}
		if len(Layer) > 0 {
			EachChilden = append(EachChilden, Layer)
		}
	}

	result = EachChilden[0][0]

	var i = 2
Loop:
	for j := range EachChilden {
		for _, node := range EachChilden[j] {
			for i < len(s) && s[i] >= 0 {
				node.Children = append(node.Children, mapNode[s[i]])
				i++
				if i >= len(s) {
					break Loop
				}
			}
			i++
		}
	}

	return result
}
