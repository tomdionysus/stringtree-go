package stringtree_test

import (
	. "github.com/tomcully/stringtree"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringTree", func() {

	Describe("Factory New", func() {

		It("should construct a tree properly", func() {
			x := NewStringTree()
			Expect(x.Root).To(BeNil())
		})

	})

	Describe("Add", func() {
		It("should form a tree correctly from a string", func() {
			tree := NewStringTree()
			tree.Add("ant", 10)

			Expect(tree.Root.Char).To(Equal('a'))
			Expect(tree.Root.Down.Char).To(Equal('n'))
			Expect(tree.Root.Down.Down.Char).To(Equal('t'))
		})

		It("should use parts of the existing tree", func() {
			tree := NewStringTree()
			tree.Add("ant", 10)
			tree.Add("ants", 11)

			Expect(tree.Root.Char).To(Equal('a'))
			Expect(tree.Root.Left).To(BeNil())
			Expect(tree.Root.Right).To(BeNil())
			Expect(tree.Root.Down.Char).To(Equal('n'))
			Expect(tree.Root.Down.Left).To(BeNil())
			Expect(tree.Root.Down.Right).To(BeNil())
			Expect(tree.Root.Down.Down.Char).To(Equal('t'))
			Expect(tree.Root.Down.Down.Left).To(BeNil())
			Expect(tree.Root.Down.Down.Right).To(BeNil())
			Expect(tree.Root.Down.Down.Down.Char).To(Equal('s'))
		})

		It("should assign value to last node", func() {
			tree := NewStringTree()
			tree.Add("ant", 10)
			Expect(tree.Root.Down.Down.Value).To(Equal(10))
		})

		XIt("should not overwrite values on path", func() {
			tree := NewStringTree()
			tree.Add("ant", 10)
			tree.Add("antop", 20)
			Expect(tree.Root.Down.Down.Value).To(Equal(10))
			Expect(tree.Root.Down.Down.Down.Down.Value).To(Equal(20))
		})
	})
})
