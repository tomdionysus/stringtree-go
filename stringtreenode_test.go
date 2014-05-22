package stringtree_test

import (
	. "github.com/tomcully/stringtree-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringTreeNode", func() {

	Describe("Factory New", func() {

		It("should construct a node properly", func() {
			x := NewStringTreeNode('t', 2)
			Expect(x.Char).To(Equal('t'))
			Expect(x.Value).To(Equal(2))
			Expect(x.Left).To(BeNil())
			Expect(x.Right).To(BeNil())
		})

	})

	Describe("Add", func() {
		var root *StringTreeNode

		BeforeEach(func() {
			root = NewStringTreeNode('t', 2)
		})

		It("should return the new node", func() {
			a := root.Add('a', 10)

			Expect(a).NotTo(BeNil())

			Expect(a.Char).To(Equal('a'))
			Expect(a.Value).To(Equal(10))
		})

		It("should add a node left if rune is less", func() {
			root.Add('a', 10)

			Expect(root.Left).NotTo(BeNil())
			Expect(root.Right).To(BeNil())

			Expect(root.Left.Char).To(Equal('a'))
			Expect(root.Left.Value).To(Equal(10))
		})

		It("should replace value if rune is equal", func() {
			root.Add('t', 20)

			Expect(root.Left).To(BeNil())
			Expect(root.Right).To(BeNil())

			Expect(root.Char).To(Equal('t'))
			Expect(root.Value).To(Equal(20))
		})

		It("should add a node right if rune is greater", func() {
			root.Add('z', 40)

			Expect(root.Left).To(BeNil())
			Expect(root.Right).NotTo(BeNil())

			Expect(root.Right.Char).To(Equal('z'))
			Expect(root.Right.Value).To(Equal(40))
		})

	})

	Describe("AddNode", func() {
		var root *StringTreeNode

		BeforeEach(func() {
			root = NewStringTreeNode('t', 2)
		})

		It("should add a node left if rune is less", func() {
			root.AddNode(NewStringTreeNode('a', 10))

			Expect(root.Left).NotTo(BeNil())
			Expect(root.Right).To(BeNil())

			Expect(root.Left.Char).To(Equal('a'))
			Expect(root.Left.Value).To(Equal(10))
		})

		It("should add a node right if rune is greater", func() {
			root.AddNode(NewStringTreeNode('z', 40))

			Expect(root.Left).To(BeNil())
			Expect(root.Right).NotTo(BeNil())

			Expect(root.Right.Char).To(Equal('z'))
			Expect(root.Right.Value).To(Equal(40))
		})

		It("should not modify tree if supplied nil", func() {
			root.AddNode(nil)

			Expect(root.Left).To(BeNil())
			Expect(root.Right).To(BeNil())
			Expect(root.Char).To(Equal('t'))
			Expect(root.Value).To(Equal(2))
		})
	})

	Describe("Find", func() {
		var (
			root *StringTreeNode
			a    *StringTreeNode
			b    *StringTreeNode
			z    *StringTreeNode
		)

		BeforeEach(func() {
			root = NewStringTreeNode('t', 2)

			a = root.Add('a', 1)
			b = root.Add('b', 5)
			z = root.Add('z', 10)
		})

		It("should find a node if it exists", func() {
			var node = root.Find('b')

			Expect(node).NotTo(BeNil())

			Expect(node.Char).To(Equal('b'))
			Expect(node.Value).To(Equal(5))
		})

		It("should returnnil if it does not exist", func() {
			var node = root.Find('p')

			Expect(node).To(BeNil())
		})
	})

	Describe("Remove", func() {

		var (
			root *StringTreeNode
			a    *StringTreeNode
			b    *StringTreeNode
			z    *StringTreeNode
		)

		BeforeEach(func() {
			root = NewStringTreeNode('t', 2)

			a = root.Add('a', 1)
			b = root.Add('b', 5)
			z = root.Add('z', 10)
		})

		It("should remove a value and retain the tree", func() {
			root = root.Remove('b', nil)

			Expect(root.Char).To(Equal('t'))

			Expect(root.Left).NotTo(BeNil())
			Expect(root.Left.Char).To(Equal('a'))
			Expect(root.Left.Left).To(BeNil())
			Expect(root.Left.Right).To(BeNil())

			Expect(root.Right).NotTo(BeNil())
			Expect(root.Right.Char).To(Equal('z'))
			Expect(root.Right.Left).To(BeNil())
			Expect(root.Right.Right).To(BeNil())
		})

		It("should remove the root value and reassign root", func() {
			root = root.Remove('t', nil)

			Expect(root.Char).To(Equal('a'))
			Expect(root.Left).To(BeNil())
			Expect(root.Right.Char).To(Equal('b'))
			Expect(root.Right.Right.Char).To(Equal('z'))
		})

		It("should remove a leaf value", func() {
			root = root.Remove('z', nil)

			Expect(root.Left).NotTo(BeNil())
			Expect(root.Left.Char).To(Equal('a'))
			Expect(root.Left.Left).To(BeNil())

			Expect(root.Left.Right).NotTo(BeNil())
			Expect(root.Left.Right.Char).To(Equal('b'))

			Expect(root.Right).To(BeNil())
		})
	})
})
