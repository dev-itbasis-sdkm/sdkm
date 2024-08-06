package scripts

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"scripts", func() {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(scripts.ReadDir(".")).Should(
			gomega.HaveLen(3),
		)
	},
)
