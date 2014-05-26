package stringtree

type StringTree struct {
	Root *StringTreeNode
}

type StringTreeMatch struct {
	Offset int
	Value  interface{}
}

func NewStringTreeMatch(offset int, value interface{}) *StringTreeMatch {
	match := new(StringTreeMatch)
	match.Offset = offset
	match.Value = value
	return match
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

func (tree *StringTree) FindLastNode(s string) *StringTreeNode {
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
	return lastNode
}

func (tree *StringTree) Find(s string) interface{} {
	var node = tree.FindLastNode(s)
	if node == nil {
		return nil
	}
	return node.Value
}

func (tree *StringTree) Remove(s string) bool {
	var node *StringTreeNode = tree.FindLastNode(s)
	if node == nil {
		return false
	} else {
		node.Value = nil
		return true
	}
}

func (tree *StringTree) Scan(data string, output chan StringTreeMatch) {

	// defer func() {
	// 	close(output)
	// }

	// if data.Length || tree.Root == nil {
	// 	return
	// }

	// var node *StringTreeNode = nil
	// lastStart, offset := 0, 0

	// for {
	// 	if node == nil {
	// 		node = tree.Root
	// 		offset = lastStart + 1
	// 		lastStart = offset
	// 	}
	// 	if lastStart >= length(data) && offset == length(data) {
	// 		break
	// 	}
	// 	foundNode := node.Find(rune(data[offset]))
	// 	if foundNode == nil {
	// 		node = tree.Root
	// 		offset = lastStart + 1
	// 		lastStart = offset
	// 	} else {
	// 		if foundNode.Value != nil {
	// 			output <- *NewStringTreeMatch(lastStart, node.Value)
	// 		}
	// 		node = foundNode.Down
	// 		offset++
	// 	}
	// }
	// close(output)
}
