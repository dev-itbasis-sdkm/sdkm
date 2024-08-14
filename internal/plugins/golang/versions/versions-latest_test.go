package versions_test

import (
	"context"
	"path"

	pluginGoVersions "github.com/dev.itbasis.sdkm/internal/plugins/golang/versions"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"Latest version", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.DescribeTable(
			"different versions of answers", func(testResponseFilePath string, wantSDKVersion sdkmSDKVersion.SDKVersion) {
				var server = initFakeServer(path.Join("testdata", "all-versions", testResponseFilePath))
				defer server.Close()

				sdkVersions := pluginGoVersions.NewVersions(server.URL())

				gomega.Expect(sdkVersions.LatestVersion(context.Background())).
					To(gomega.Equal(wantSDKVersion))
			},
			ginkgo.Entry(nil, "001.html", sdkmSDKVersion.SDKVersion{ID: "1.22.5", Type: sdkmSDKVersion.TypeStable}),
			ginkgo.Entry(nil, "002.html", sdkmSDKVersion.SDKVersion{ID: "1.23.0", Type: sdkmSDKVersion.TypeStable}),
		)
	},
)
