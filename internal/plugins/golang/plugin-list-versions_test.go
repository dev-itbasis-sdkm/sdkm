package golang_test

import (
	"context"

	sdkmPluginGo "github.com/dev.itbasis.sdkm/internal/plugins/golang"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = ginkgo.Describe(
	"ListAllVersions", func() {
		defer ginkgo.GinkgoRecover()

	},
)

var _ = ginkgo.Describe(
	"ListAllVersionsByPrefix", func() {
		defer ginkgo.GinkgoRecover()

		var pluginGo sdkmPlugin.SDKMPlugin

		ginkgo.BeforeEach(
			func() {
				mockController := gomock.NewController(ginkgo.GinkgoT())

				mockSDKVersions := sdkmPlugin.NewMockSDKVersions(mockController)
				mockSDKVersions.EXPECT().
					AllVersions(gomock.Any()).
					Return(
						[]sdkmPlugin.SDKVersion{
							{ID: "1.22.5", Type: sdkmPlugin.TypeStable},
							{ID: "1.21.12", Type: sdkmPlugin.TypeStable},
							{ID: "1.23rc2", Type: sdkmPlugin.TypeUnstable},
							{ID: "1.23rc1", Type: sdkmPlugin.TypeArchived},
							{ID: "1.22.4", Type: sdkmPlugin.TypeArchived},
							{ID: "1.22.3", Type: sdkmPlugin.TypeArchived},
							{ID: "1.22.0", Type: sdkmPlugin.TypeArchived},
							{ID: "1.21.11", Type: sdkmPlugin.TypeArchived},
							{ID: "1.21.10", Type: sdkmPlugin.TypeArchived},
							{ID: "1.21.0", Type: sdkmPlugin.TypeArchived},
							{ID: "1.21rc3", Type: sdkmPlugin.TypeArchived},
							{ID: "1.20.14", Type: sdkmPlugin.TypeArchived},
							{ID: "1.19.13", Type: sdkmPlugin.TypeArchived},
							{ID: "1.19.12", Type: sdkmPlugin.TypeArchived},
							{ID: "1.18", Type: sdkmPlugin.TypeArchived},
							{ID: "1.18.10", Type: sdkmPlugin.TypeArchived},
							{ID: "1.4beta1", Type: sdkmPlugin.TypeArchived},
							{ID: "1.3rc1", Type: sdkmPlugin.TypeArchived},
						},
					).MaxTimes(1)

				mockBasePlugin := sdkmPlugin.NewMockBasePlugin(mockController)
				mockBasePlugin.EXPECT().GetSDKDir().Return("").AnyTimes()
				mockBasePlugin.EXPECT().HasInstalled("go", gomock.Any()).Return(false).AnyTimes()

				pluginGo = sdkmPluginGo.GetPlugin().
					WithBasePlugin(mockBasePlugin).
					WithVersions(mockSDKVersions)
			},
		)

		ginkgo.It(
			"", func() {
				gomega.Expect(pluginGo.ListAllVersions(context.Background())).
					Should(
						gomega.SatisfyAll(
							gomega.HaveLen(18),

							gomega.ContainElements(
								sdkmPlugin.SDKVersion{ID: "1.23rc2", Type: sdkmPlugin.TypeUnstable},
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

		ginkgo.When(
			"By Prefix", func() {
				ginkgo.It(
					"empty prefix", func() {
						gomega.Expect(pluginGo.ListAllVersionsByPrefix(context.Background(), "")).
							Should(
								gomega.SatisfyAll(
									gomega.HaveLen(18),

									gomega.ContainElements(
										sdkmPlugin.SDKVersion{ID: "1.23rc2", Type: sdkmPlugin.TypeUnstable},
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

				ginkgo.DescribeTable(
					"success", func(prefix string, wantCount int, wantSDKVersions []sdkmPlugin.SDKVersion) {
						gomega.Expect(pluginGo.ListAllVersionsByPrefix(context.Background(), prefix)).
							Should(
								gomega.SatisfyAll(
									gomega.HaveLen(wantCount),
									gomega.ContainElements(wantSDKVersions),
								),
							)
					},
					ginkgo.Entry(
						nil, "1.23", 2, []sdkmPlugin.SDKVersion{
							{ID: "1.23rc2", Type: sdkmPlugin.TypeUnstable},
							{ID: "1.23rc1", Type: sdkmPlugin.TypeArchived},
						},
					),
					ginkgo.Entry(
						nil, "1.21", 5, []sdkmPlugin.SDKVersion{
							{ID: "1.21.12", Type: sdkmPlugin.TypeStable},
							{ID: "1.21.11", Type: sdkmPlugin.TypeArchived},
							{ID: "1.21.0", Type: sdkmPlugin.TypeArchived},
							{ID: "1.21rc3", Type: sdkmPlugin.TypeArchived},
						},
					),
				)
			},
		)
	},
)
