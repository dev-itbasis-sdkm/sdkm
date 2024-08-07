package versions_test

import (
	"context"

	pluginGoVersions "github.com/dev.itbasis.sdkm/internal/plugins/golang/versions"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.FDescribe(
	"All Versions", func() {
		defer ginkgo.GinkgoRecover()

		var server = testServer() //nolint:ginkgolinter // TODO
		defer server.Close()

		gomega.Expect(pluginGoVersions.NewVersions(server.URL()).AllVersions(context.Background())).
			To(
				gomega.SatisfyAll(
					gomega.HaveLen(292),

					gomega.ContainElements(
						sdkmSDKVersion.SDKVersion{ID: "1.23rc2", Type: sdkmSDKVersion.TypeUnstable},
						sdkmSDKVersion.SDKVersion{ID: "1.23rc1", Type: sdkmSDKVersion.TypeArchived},
						sdkmSDKVersion.SDKVersion{ID: "1.22.5", Type: sdkmSDKVersion.TypeStable},
						sdkmSDKVersion.SDKVersion{ID: "1.22.0", Type: sdkmSDKVersion.TypeArchived},
						sdkmSDKVersion.SDKVersion{ID: "1.18", Type: sdkmSDKVersion.TypeArchived},
						sdkmSDKVersion.SDKVersion{ID: "1.18.10", Type: sdkmSDKVersion.TypeArchived},
						sdkmSDKVersion.SDKVersion{ID: "1.4beta1", Type: sdkmSDKVersion.TypeArchived},
						sdkmSDKVersion.SDKVersion{ID: "1.3rc1", Type: sdkmSDKVersion.TypeArchived},
					),
				),
			)
	},
)
