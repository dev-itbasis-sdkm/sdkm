package sdkversion_test

import (
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"GetCacheFilePath", func() {
		gomega.Expect(sdkmSDKVersion.GetCacheFilePath("pn")).
			To(gomega.HaveSuffix("/.cache/pn.json"))
	},
)
