stringtree-go
=============

StringTree in Golang.

## Global Methods

* ```*StringTree NewStringTree()``` - Create a new empty StringTree
* ```*StringTreeNode NewStringTreeNode(c rune, v int)``` - Create a new node with the specified rune and value

## Types

### StringTree

**Exported Fields**

* ```Root *StringTreeNode``` - The root node, or nil if the tree is empty.

**Methods**

* ```void Add(s string, value interface{})``` - Add a string and object to the tree
* ```interface{} Find(s string)``` - Find a string in the tree and return its object, or nil if not found

### StringTreeNode

**Exported Fields**

* ```Char rune``` - The rune value of this node
* ```Value interface{}``` - The object 
* ```Left *StringTreeNode``` - The left branch of this node
* ```Right *StringTreeNode``` - The right branch of this node
* ```Down *StringTreeNode``` - The down (or forward) root of the next rune

**Methods**

* ```*StringTreeNode Add(c rune, v int)``` - Add a new leaf with the specified values and return the new node
* ```void AddNode(newNode *StringTreeNode)``` - Add an existing node to this tree
* ```*StringTreeNode Find(c rune)``` - Return the node with the specified rune, or nil if not found.
* ```*StringTreeNode Remove(c rune, parent *StringTreeNode)``` - Remove the node with the specified rune from the tree and return the new tree root.
* ```void Iterate(f func(char rune, i interface{}))``` - Call the function f for each rune/value in the tree in ascending order

## Testing

Stringtree-Go uses [ginkgo](http://onsi.github.io/ginkgo) and [gomega](http://onsi.github.io/gomega). Run the test suite as follows:

```
	cd stringtree
	go test
```