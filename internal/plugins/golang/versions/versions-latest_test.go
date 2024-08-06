package versions_test

import (
	"context"

	pluginGoVersions "github.com/dev.itbasis.sdkm/internal/plugins/golang/versions"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"Latest version", func() {
		defer ginkgo.GinkgoRecover()

		var server = testServer()
		defer server.Close()

		sdkVersions := pluginGoVersions.NewVersions(server.URL())

		gomega.Expect(sdkVersions.LatestVersion(context.Background())).
			Should(gomega.Equal(sdkmPlugin.SDKVersion{ID: "1.22.5", Type: sdkmPlugin.TypeStable}))
	},
)
