package stringtree_test

import (
	. "github.com/tomcully/stringtree-go"

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

		It("should overwrite value in existing last node", func() {
			tree := NewStringTree()
			tree.Add("ant", 10)
			tree.Add("ant", 40)
			Expect(tree.Root.Down.Down.Value).To(Equal(40))
		})

		It("should not overwrite values on path", func() {
			tree := NewStringTree()
			tree.Add("ant", 10)
			tree.Add("antop", 20)
			Expect(tree.Root.Down.Down.Value).To(Equal(10))
			Expect(tree.Root.Down.Down.Down.Down.Value).To(Equal(20))
		})

		It("should assign value even if path continues", func() {
			tree := NewStringTree()
			tree.Add("antop", 20)
			tree.Add("ant", 10)
			Expect(tree.Root.Down.Down.Value).To(Equal(10))
			Expect(tree.Root.Down.Down.Down.Down.Value).To(Equal(20))
		})
	})

	Describe("FindLastNode", func() {

		var tree *StringTree

		BeforeEach(func() {
			tree = NewStringTree()
			tree.Add("aardvark", 10)
			tree.Add("anteater", 20)
			tree.Add("ant", 30)
			tree.Add("tom", 40)
			tree.Add("tom cully", 50)
			tree.Add("whale", 60)
			tree.Add("zebra", 70)
		})

		It("should return last StringTreeNode for strings in tree", func() {
			Expect(tree.FindLastNode("whale").Char).To(Equal('e'))
			Expect(tree.FindLastNode("whale").Value).To(Equal(60))
			Expect(tree.FindLastNode("tom").Value).To(Equal(40))
			Expect(tree.FindLastNode("anteater").Value).To(Equal(20))
			Expect(tree.FindLastNode("zebra").Value).To(Equal(70))
			Expect(tree.FindLastNode("aardvark").Value).To(Equal(10))
			Expect(tree.FindLastNode("tom cully").Value).To(Equal(50))
		})

		It("should return lastnode for partial strings not in tree", func() {
			Expect(tree.FindLastNode("whal").Char).To(Equal('l'))
			Expect(tree.FindLastNode("aneater").Char).To(Equal('n'))
			Expect(tree.FindLastNode("ebra")).To(BeNil())
			Expect(tree.FindLastNode("tomcully").Char).To(Equal('m'))
			Expect(tree.FindLastNode("o")).To(BeNil())
			Expect(tree.FindLastNode("")).To(BeNil())
		})

		XIt("should return nil for any string when tree is empty", func() {
			tree = NewStringTree()

			Expect(tree.Find("whale")).To(BeNil())
			Expect(tree.Find("")).To(BeNil())
		})
	})

	Describe("Find", func() {

		var tree *StringTree

		BeforeEach(func() {
			tree = NewStringTree()
			tree.Add("aardvark", 10)
			tree.Add("anteater", 20)
			tree.Add("ant", 30)
			tree.Add("tom", 40)
			tree.Add("tom cully", 50)
			tree.Add("whale", 60)
			tree.Add("zebra", 70)
		})

		It("should return values for strings in tree", func() {
			Expect(tree.Find("whale")).To(Equal(60))
			Expect(tree.Find("tom")).To(Equal(40))
			Expect(tree.Find("anteater")).To(Equal(20))
			Expect(tree.Find("zebra")).To(Equal(70))
			Expect(tree.Find("aardvark")).To(Equal(10))
			Expect(tree.Find("tom cully")).To(Equal(50))
		})

		It("should return nil for strings not in tree", func() {
			Expect(tree.Find("whal")).To(BeNil())
			Expect(tree.Find("aneater")).To(BeNil())
			Expect(tree.Find("ebra")).To(BeNil())
			Expect(tree.Find("tomcully")).To(BeNil())
			Expect(tree.Find("o")).To(BeNil())
			Expect(tree.Find("")).To(BeNil())
		})

		It("should return nil for any string when tree is empty", func() {
			tree = NewStringTree()

			Expect(tree.Find("whale")).To(BeNil())
			Expect(tree.Find("")).To(BeNil())
		})
	})

	Describe("Remove", func() {

		var tree *StringTree

		BeforeEach(func() {
			tree = NewStringTree()
			tree.Add("aardvark", 10)
			tree.Add("anteater", 20)
			tree.Add("ant", 30)
			tree.Add("tom", 40)
			tree.Add("tom cully", 50)
			tree.Add("whale", 60)
			tree.Add("zebra", 70)
		})

		It("should a remove string", func() {
			tree.Remove("whale")
			Expect(tree.Find("whale")).To(BeNil())
		})

		It("should not break tree when a string is removed", func() {
			tree.Remove("whale")

			Expect(tree.Find("tom")).To(Equal(40))
			Expect(tree.Find("anteater")).To(Equal(20))
			Expect(tree.Find("zebra")).To(Equal(70))
			Expect(tree.Find("aardvark")).To(Equal(10))
			Expect(tree.Find("tom cully")).To(Equal(50))
			Expect(tree.Find("whale")).To(BeNil())
		})

		It("should prune tree back to last value on remove", func() {

		})
	})

	Describe("Scan", func() {

		var tree *StringTree

		BeforeEach(func() {
			tree = NewStringTree()
			tree.Add("one", 10)
			tree.Add("two", 20)
			tree.Add("three", 30)
			tree.Add("four", 40)
			tree.Add("five", 50)
			tree.Add("six", 60)
			tree.Add("seven", 70)
			tree.Add("fourteen", 140)
		})

		It("should find basic strings with correct offsets", func() {

			output := make(chan StringTreeMatch)

			tree.Scan("onetwothree", output)

			// Channel output should be [0,10]
			// Channel output should be [3,20]
			// Channel output should be [6,30]
			// Channel should be closed
		})
	})
})
