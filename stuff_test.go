package counterfeiterrepro

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jfmyers9/counterfeiterrepro/counterfeiterreprofakes"
)

var _ = Describe("test", func() {
	var thingDoer *counterfeiterreprofakes.FakeThingDoer

	It("works", func() {
		thingDoer = &counterfeiterreprofakes.FakeThingDoer{}
		thingDoer.DoThingReturn(nil)
		Expect(true).To(BeTrue())
	})
})
