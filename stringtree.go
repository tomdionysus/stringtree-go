package stringtree

type StringTreeNode struct {
	Char  rune
	Value int
	Left  *StringTreeNode
	Right *StringTreeNode
	Down  *StringTreeNode
}

func NewStringTreeNode(c rune, v int) *StringTreeNode {
	node := new(StringTreeNode)
	node.Char = c
	node.Value = v
	node.Left = nil
	node.Right = nil
	return node
}

func (node *StringTreeNode) Add(c rune, v int) *StringTreeNode {
	var newNode *StringTreeNode
	if c <= node.Char {
		if node.Left == nil {
			node.Left = new(StringTreeNode)
			node.Left.Char = c
			node.Left.Value = v
			newNode = node.Left
		} else {
			newNode = node.Left.Add(c, v)
		}
	} else {
		if node.Right == nil {
			node.Right = new(StringTreeNode)
			node.Right.Char = c
			node.Right.Value = v
			newNode = node.Right
		} else {
			newNode = node.Right.Add(c, v)
		}
	}
	return newNode
}

func (node *StringTreeNode) Find(c rune) *StringTreeNode {
	if c == node.Char {
		return node
	}
	if c <= node.Char && node.Left != nil {
		return node.Left.Find(c)
	} else if c > node.Char && node.Right != nil {
		return node.Right.Find(c)
	}
	return nil
}

func (node *StringTreeNode) Iterate(f func(c rune, i int)) {
	if node.Left != nil {
		node.Left.Iterate(f)
	}
	f(node.Char, node.Value)
	if node.Right != nil {
		node.Right.Iterate(f)
	}
}

type StringTree struct {
	root StringTreeNode
}
