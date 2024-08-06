package plugin_test

import (
	"github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"GetCacheFilePath", func() {
		gomega.Expect(plugin.GetCacheFilePath("pn")).
			Should(gomega.HaveSuffix("/.cache/pn.json"))
	},
)
