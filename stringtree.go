package stringtree

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
