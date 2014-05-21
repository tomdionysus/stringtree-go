package stringtree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStringtree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringtree Suite")
}
