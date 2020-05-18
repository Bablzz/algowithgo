package algo

type KV interface {
	Less(KV) bool
	Equal(KV) bool
}

type Node struct {
	KeyValue     KV
	BalanceValue int
	Nodes        [2]*Node
}

type IntegerKey int

func (k IntegerKey) Less(k1 KV) bool {
	return k < k1.(IntegerKey)
}
func (k IntegerKey) Equal(k1 KV) bool {
	return k == k1.(IntegerKey)
}

func opposite(nodeValue int) int {
	return 1 - nodeValue
}

func InsertNode(treeNode **Node, key KV) {
	*treeNode, _ = insertRNode(*treeNode, key)
}

func RemoveNode(treeNode **Node, key KV) {
	*treeNode, _ = removeRNode(*treeNode, key)
}

func singleRotation(root *Node, nodeValue int) *Node {
	var saveNode *Node
	saveNode = root.Nodes[opposite(nodeValue)]
	root.Nodes[opposite(nodeValue)] = saveNode.Nodes[nodeValue]
	saveNode.Nodes[nodeValue] = root
	return saveNode
}

func doubleRotation(rootNode *Node, nodeValue int) *Node {
	var saveNode *Node
	saveNode =
		rootNode.Nodes[opposite(nodeValue)].Nodes[nodeValue]
	rootNode.Nodes[opposite(nodeValue)].Nodes[nodeValue] =
		saveNode.Nodes[opposite(nodeValue)]
	saveNode.Nodes[opposite(nodeValue)] =
		rootNode.Nodes[opposite(nodeValue)]
	rootNode.Nodes[opposite(nodeValue)] = saveNode
	saveNode = rootNode.Nodes[opposite(nodeValue)]
	rootNode.Nodes[opposite(nodeValue)] = saveNode.Nodes[nodeValue]
	saveNode.Nodes[nodeValue] = rootNode
	return saveNode
}

func adjustBalance(rootNode *Node, nodeValue int, balanceValue int) {
	var (
		node    *Node
		oppNode *Node
	)
	node = rootNode.Nodes[nodeValue]
	oppNode = node.Nodes[opposite(balanceValue)]
	switch oppNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

func BalanceTree(rootNode *Node, nodeValue int) *Node {
	var node *Node
	node = rootNode.Nodes[nodeValue]
	var balance int
	balance = 2*nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, opposite(nodeValue))
	}
	adjustBalance(rootNode, nodeValue, balance)
	return doubleRotation(rootNode, opposite(nodeValue))
}

func insertRNode(rootNode *Node, key KV) (*Node, bool) {
	var (
		dir  int
		done bool
	)

	if rootNode == nil {
		return &Node{KeyValue: key}, false
	}

	dir = 0
	if rootNode.KeyValue.Less(key) {
		dir = 1
	}

	rootNode.Nodes[dir], done = insertRNode(rootNode.Nodes[dir],
		key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}
	return BalanceTree(rootNode, dir), true
}

func removeRNode(rootNode *Node, key KV) (*Node, bool) {
	if rootNode == nil {
		return nil, false
	}
	if rootNode.KeyValue.Equal(key) {
		switch {
		case rootNode.Nodes[0] == nil:
			return rootNode.Nodes[1], false
		case rootNode.Nodes[1] == nil:
			return rootNode.Nodes[0], false
		}
		var heirNode *Node
		heirNode = rootNode.Nodes[0]
		for heirNode.Nodes[1] != nil {
			heirNode = heirNode.Nodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		key = heirNode.KeyValue
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.Less(key) {
		dir = 1
	}
	var done bool
	rootNode.Nodes[dir], done = removeRNode(rootNode.Nodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (1 - 2*dir)
	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}
	return removeBalance(rootNode, dir)
}

func removeBalance(rootNode *Node, nodeValue int) (*Node, bool) {
	var node *Node
	node = rootNode.Nodes[opposite(nodeValue)]
	var balance int
	balance = 2*nodeValue - 1
	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, nodeValue), false
	case balance:
		adjustBalance(rootNode, opposite(nodeValue), -balance)
		return doubleRotation(rootNode, nodeValue), false
	}
	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return singleRotation(rootNode, nodeValue), true
}
