Stringtree-Go
=============

StringTree in Golang.

## Globals

StringTreeNode NewStringTreeNode(c rune, v int) - Create A StringTreeNode with the specified rune and value

## Types

### StringTree

Exported Fields
* Root

Methods
* Add

### StringTreeNode

Exported Fields
* Char
* Value
* Left
* Right
* Down

Methods
* Add(c rune, v int) - Add a new leaf with the specified values
* Find(c rune) - Return the StringTreeNode with the specified rune, or nil if not found.

## Testing

Stringtree-Go uses [ginkgo](http://onsi.github.io/ginkgo) and [gomega](http://onsi.github.io/gomega).

```
	cd stringtree
	go test
```