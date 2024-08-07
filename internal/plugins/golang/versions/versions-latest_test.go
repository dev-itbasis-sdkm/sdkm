package versions_test

import (
	"context"

	pluginGoVersions "github.com/dev.itbasis.sdkm/internal/plugins/golang/versions"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"Latest version", func() {
		defer ginkgo.GinkgoRecover()

		var server = testServer() //nolint:ginkgolinter // TODO
		defer server.Close()

		gomega.Expect(
			pluginGoVersions.NewVersions(server.URL()).
				LatestVersion(context.Background()),
		).
			To(gomega.Equal(sdkmSDKVersion.SDKVersion{ID: "1.22.5", Type: sdkmSDKVersion.TypeStable}))
	},
)
