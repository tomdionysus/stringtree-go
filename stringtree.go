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

	var newNode **StringTreeNode
	if char < node.Char {
		newNode = &(node.Left)
	} else {
		newNode = &(node.Right)
	}
	if (*newNode) == nil {
		*newNode = NewStringTreeNode(char, value)
		return *newNode
	}
	return (*newNode).Add(char, value)
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

type StringTree struct {
	Root *StringTreeNode
}

func NewStringTree() *StringTree {
	tree := new(StringTree)
	tree.Root = nil
	return tree
}

func (tree *StringTree) Add(s string, value interface{}) {
	var node, lastNode **StringTreeNode
	node = &tree.Root
	for _, char := range s {
		lastNode = node
		if *node == nil {
			*node = NewStringTreeNode(char, nil)
			node = &(*node).Down
		} else {
			node = &((*node).Add(char, nil)).Down
		}
	}
	(*lastNode).Value = value
}

func (tree *StringTree) Find(s string) interface{} {
	if tree.Root == nil {
		return nil
	}
	var node *StringTreeNode = tree.Root
	var lastNode *StringTreeNode
	for _, char := range s {
		node = node.Find(char)
		if node == nil {
			return nil
		}
		lastNode = node
		node = node.Down
	}
	if lastNode == nil {
		return nil
	}
	return lastNode.Value
}
