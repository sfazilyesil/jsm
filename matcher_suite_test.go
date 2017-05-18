package jsm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestJSONSchemaMatcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JSONSchemaMatcher Suite")
}
