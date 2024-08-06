package golang_test

import (
	"context"
	"fmt"

	sdkmPluginGo "github.com/dev.itbasis.sdkm/internal/plugins/golang"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"go.uber.org/mock/gomock"
)

var _ = ginkgo.Describe(
	"Plugin Latest Version", func() {
		defer ginkgo.GinkgoRecover()

		var pluginGo sdkmPlugin.SDKMPlugin

		ginkgo.BeforeEach(
			func() {
				mockController := gomock.NewController(ginkgo.GinkgoT())

				mockSDKVersions := sdkmPlugin.NewMockSDKVersions(mockController)
				mockSDKVersions.EXPECT().LatestVersion(gomock.Any()).Return(sdkmPlugin.SDKVersion{ID: "1.22.5", Type: sdkmPlugin.TypeStable})

				mockBasePlugin := sdkmPlugin.NewMockBasePlugin(mockController)
				mockBasePlugin.EXPECT().GetSDKDir().Return("").AnyTimes()

				pluginGo = sdkmPluginGo.GetPlugin().
					WithBasePlugin(mockBasePlugin).
					WithVersions(mockSDKVersions)

			},
		)

		ginkgo.DescribeTable(
			"LatestVersion", func(wantSDKVersion sdkmPlugin.SDKVersion) {
				gomega.Expect(pluginGo.LatestVersion(context.Background())).
					Should(
						gomega.HaveValue(
							gstruct.MatchFields(
								gstruct.IgnoreExtras, gstruct.Fields{
									"ID":   gomega.Equal(wantSDKVersion.ID),
									"Type": gomega.Equal(wantSDKVersion.Type),
								},
							),
						),
					)
			},
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{ID: "1.22.5", Type: sdkmPlugin.TypeStable}),
		)
	},
)

var _ = ginkgo.Describe(
	"LatestVersionByPrefix", func() {
		defer ginkgo.GinkgoRecover()

		var pluginGo sdkmPlugin.SDKMPlugin

		ginkgo.BeforeEach(
			func() {
				mockController := gomock.NewController(ginkgo.GinkgoT())

				mockSDKVersions := sdkmPlugin.NewMockSDKVersions(mockController)
				mockSDKVersions.EXPECT().
					LatestVersion(gomock.Any()).
					Return(sdkmPlugin.SDKVersion{ID: "1.22.5", Type: sdkmPlugin.TypeStable}).
					MaxTimes(1)
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
							{ID: "1.21.11", Type: sdkmPlugin.TypeArchived},
							{ID: "1.20.14", Type: sdkmPlugin.TypeArchived},
							{ID: "1.19.13", Type: sdkmPlugin.TypeArchived},
							{ID: "1.19.12", Type: sdkmPlugin.TypeArchived},
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

		ginkgo.DescribeTable(
			"success", func(prefix string, wantSDKVersion sdkmPlugin.SDKVersion) {
				gomega.Expect(pluginGo.LatestVersionByPrefix(context.Background(), prefix)).
					Should(
						gomega.HaveValue(
							gstruct.MatchFields(
								gstruct.IgnoreExtras, gstruct.Fields{
									"ID":   gomega.Equal(wantSDKVersion.ID),
									"Type": gomega.Equal(wantSDKVersion.Type),
								},
							),
						),
					)
			},
			ginkgo.Entry("empty prefix", "", sdkmPlugin.SDKVersion{ID: "1.22.5", Type: sdkmPlugin.TypeStable}),
			ginkgo.Entry(nil, "1.23", sdkmPlugin.SDKVersion{ID: "1.23rc2", Type: sdkmPlugin.TypeUnstable}),
			ginkgo.Entry(nil, "1.22", sdkmPlugin.SDKVersion{ID: "1.22.5", Type: sdkmPlugin.TypeStable}),
			ginkgo.Entry(nil, "1.21", sdkmPlugin.SDKVersion{ID: "1.21.12", Type: sdkmPlugin.TypeStable}),
			ginkgo.Entry(nil, "1.20", sdkmPlugin.SDKVersion{ID: "1.20.14", Type: sdkmPlugin.TypeArchived}),
			ginkgo.Entry(nil, "1.19", sdkmPlugin.SDKVersion{ID: "1.19.13", Type: sdkmPlugin.TypeArchived}),
		)

		ginkgo.DescribeTable(
			"fail", func(prefix string) {
				gomega.Expect(pluginGo.LatestVersionByPrefix(context.Background(), prefix)).Error().Should(
					gomega.MatchError(fmt.Sprintf("version by prefix %s: SDK version not found", prefix)),
				)
			},
			ginkgo.Entry("", "1.24"),
		)
	},
)
