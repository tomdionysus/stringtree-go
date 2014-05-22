package stringtree

type StringTreeNode struct {
	Char  rune
	Value interface{}
	Left  *StringTreeNode
	Right *StringTreeNode
	Down  *StringTreeNode
}

func NewStringTreeNode(char rune, value interface{}) *StringTreeNode {
	node := new(StringTreeNode)
	node.Char = char
	node.Value = value
	node.Left = nil
	node.Right = nil
	return node
}

func (node *StringTreeNode) Add(char rune, value interface{}) *StringTreeNode {
	if node.Char == char {
		if value != nil {
			node.Value = value
		}
		return node
	}

	var nodePtr **StringTreeNode
	if char < node.Char {
		nodePtr = &(node.Left)
	} else {
		nodePtr = &(node.Right)
	}
	if *nodePtr == nil {
		*nodePtr = NewStringTreeNode(char, value)
		return *nodePtr
	}
	return (*nodePtr).Add(char, value)
}

func (node *StringTreeNode) AddNode(newNode *StringTreeNode) {

	var nodePtr **StringTreeNode
	if newNode.Char < node.Char {
		nodePtr = &(node.Left)
	} else {
		nodePtr = &(node.Right)
	}
	if *nodePtr == nil {
		*nodePtr = newNode
	} else {
		(*nodePtr).AddNode(newNode)
	}
}

func (node *StringTreeNode) Find(char rune) *StringTreeNode {
	if char == node.Char {
		return node
	}
	if char <= node.Char && node.Left != nil {
		return node.Left.Find(char)
	} else if char > node.Char && node.Right != nil {
		return node.Right.Find(char)
	}
	return nil
}

func (node *StringTreeNode) Iterate(f func(char rune, i interface{})) {
	if node.Left != nil {
		node.Left.Iterate(f)
	}
	f(node.Char, node.Value)
	if node.Right != nil {
		node.Right.Iterate(f)
	}
}
