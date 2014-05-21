package stringtree_test

import (
	. "github.com/tomcully/stringtree"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("stringtree", func() {

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

		It("should add a node left if rune value is less", func() {

			root.Add('a', 10)

			Expect(root.Left).NotTo(BeNil())
			Expect(root.Right).To(BeNil())

			Expect(root.Left.Char).To(Equal('a'))
			Expect(root.Left.Value).To(Equal(10))

		})

		It("should add a node left if rune value is equal", func() {

			root.Add('t', 20)

			Expect(root.Left).NotTo(BeNil())
			Expect(root.Right).To(BeNil())

			Expect(root.Left.Char).To(Equal('t'))
			Expect(root.Left.Value).To(Equal(20))

		})

		It("should add a node right if rune value is greater", func() {

			root.Add('z', 40)

			Expect(root.Left).To(BeNil())
			Expect(root.Right).NotTo(BeNil())

			Expect(root.Right.Char).To(Equal('z'))
			Expect(root.Right.Value).To(Equal(40))

		})

	})
})
