stringtree-go
=============

StringTree in Golang.

StringTree is a fast forward pattern matcher, suited to:

* Tokenization
* Virus Signature Scanning
* Partial matching, e.g. autocomplete suggestions

# Implementation

StringTree is a n-dimensional *trie* holding a map of strings to objects (```interface{}``` in Go), where:

* The first letters of each string are represented by a binary tree,
* The second letters of all strings starting with a specific letter are represented by a binary tree that 'descends' from the node containing that letter in the first tree,
* The third letters of all strings starting with a specifc two characters are represented in a tree descending from the node of the second letter, and so on.

During scanning, StringTree starts on each letter with a search in the first tree, on finding a match proceeds to the next letter in the descending tree, and so on. When a value node is found (the end of an added string with an object ref) the object is added to the output list, with its offset in the data. Scanning then begins again from the next letter until all data is exhausted, producing a list of objects and offsets for each matched string in the data.

# Package

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

		cd stringtree
		go test
