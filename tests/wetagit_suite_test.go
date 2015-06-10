package tests

import (
	_ "github.com/drborges/wetagit/tests/controllers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWetagit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wetagit Suite")
}
