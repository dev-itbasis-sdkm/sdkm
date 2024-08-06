package versions_test

import (
	"context"

	pluginGoVersions "github.com/dev.itbasis.sdkm/internal/plugins/golang/versions"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"All Versions", func() {
		defer ginkgo.GinkgoRecover()

		var server = testServer()
		defer server.Close()

		gomega.Expect(pluginGoVersions.NewVersions(server.URL()).AllVersions(context.Background())).
			Should(
				gomega.SatisfyAll(
					gomega.HaveLen(292),

					gomega.ContainElements(
						sdkmPlugin.SDKVersion{ID: "1.23rc2", Type: sdkmPlugin.TypeUnstable},
						sdkmPlugin.SDKVersion{ID: "1.23rc1", Type: sdkmPlugin.TypeArchived},
						sdkmPlugin.SDKVersion{ID: "1.22.5", Type: sdkmPlugin.TypeStable},
						sdkmPlugin.SDKVersion{ID: "1.22.0", Type: sdkmPlugin.TypeArchived},
						sdkmPlugin.SDKVersion{ID: "1.18", Type: sdkmPlugin.TypeArchived},
						sdkmPlugin.SDKVersion{ID: "1.18.10", Type: sdkmPlugin.TypeArchived},
						sdkmPlugin.SDKVersion{ID: "1.4beta1", Type: sdkmPlugin.TypeArchived},
						sdkmPlugin.SDKVersion{ID: "1.3rc1", Type: sdkmPlugin.TypeArchived},
					),
				),
			)
	},
)
