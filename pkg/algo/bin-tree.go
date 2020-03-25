package algo

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (n *Node) Insert(val int) {
	if n.Value > val {
		if n.Left == nil {
			n.Left = &Node{
				Value: val,
			}
		}
		n.Left.Insert(val)
	}

	if n.Value < val {
		if n.Right == nil {
			n.Right = &Node{
				Value: val,
			}
		}
		n.Right.Insert(val)
	}
}

//func (n *Node) Walk(nd *Node) int {
//	if n.Right == nil  && n.Left == nil {
//		fmt.Printf("Node %#v", n)
//		return n.Value
//	}
//	if n.Right == nil  && n.Left != nil {
//		n.Walk(nd.Left)
//	}
//	if n.Right != nil  && n.Left == nil {
//		n.Walk(nd.Right)
//	}
//	return 0
//}

//func (n *Node) Delete(val int) {
//	fmt.Print("delete work")
//	if n.Left == nil && n.Right == nil {
//		fmt.Printf("Nothing to delete")
//		return
//	}
//    // case 1. have no right child
//    if val < n.Value {
//		if n.Left != nil {
//			if n.Left.Right == nil && n.Left.Left != nil {
//				n.Left.Value =  n.Left.Left.Value
//				n.Left.Left = nil
//			}
//		}
//	}
//}
