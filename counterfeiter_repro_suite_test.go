package counterfeiterrepro_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCounterfeiterRepro(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CounterfeiterRepro Suite")
}
