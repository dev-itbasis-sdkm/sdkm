package golang_test

import (
	"context"
	"path"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmPluginGo "github.com/dev.itbasis.sdkm/internal/plugins/golang"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"go.uber.org/mock/gomock"
)

var _ = ginkgo.Describe(
	"Current", func() {
		type testData struct {
			sdkInstalled  bool
			dir           string
			wantID        string
			wantInstalled bool
		}

		var (
			mockBasePlugin *sdkmPlugin.MockBasePlugin
			pluginGo       sdkmPlugin.SDKMPlugin
		)

		ginkgo.BeforeEach(
			func() {
				mockController := gomock.NewController(ginkgo.GinkgoT())

				mockSDKVersions := sdkmSDKVersion.NewMockSDKVersions(mockController)
				mockSDKVersions.EXPECT().AllVersions(gomock.Any()).Return(
					[]sdkmSDKVersion.SDKVersion{
						{ID: "1.22.5"},
						{ID: "1.22.4"},
						{ID: "1.22.3"},
						{ID: "1.23rc2"},
						{ID: "1.21.12"},
						{ID: "1.21.11"},
						{ID: "1.23rc1"},
						{ID: "1.19.12"},
					},
				)

				mockBasePlugin = sdkmPlugin.NewMockBasePlugin(mockController)
				mockBasePlugin.EXPECT().GetSDKDir().Return("").AnyTimes()

				pluginGo = sdkmPluginGo.GetPlugin().
					WithBasePlugin(mockBasePlugin).
					WithVersions(mockSDKVersions)
			},
		)

		ginkgo.DescribeTable(
			"success", func(testData testData) {
				mockBasePlugin.EXPECT().HasInstalled("go", testData.wantID).
					Return(testData.wantInstalled).
					MaxTimes(2)

				baseDir := path.Join(sdkmOs.Pwd(), "testdata/current", testData.dir)
				gomega.Expect(baseDir).To(gomega.BeADirectory())

				gomega.Expect(pluginGo.Current(context.Background(), baseDir)).
					To(
						gomega.HaveValue(
							gstruct.MatchFields(
								gstruct.IgnoreExtras, gstruct.Fields{
									"ID":        gomega.Equal(testData.wantID),
									"Installed": gomega.Equal(testData.wantInstalled),
								},
							),
						),
					)
			},
			ginkgo.Entry(nil, testData{dir: "001", wantID: "1.21.12"}),
			ginkgo.Entry(nil, testData{dir: "002", wantID: "1.22.5"}),
			ginkgo.Entry(nil, testData{dir: "002", sdkInstalled: true, wantID: "1.22.5", wantInstalled: true}),
			ginkgo.Entry(nil, testData{dir: "003", sdkInstalled: true, wantID: "1.22.5", wantInstalled: true}),
			ginkgo.Entry(nil, testData{dir: "003", wantID: "1.22.5"}),
			ginkgo.Entry(nil, testData{dir: "004", sdkInstalled: true, wantID: "1.22.3", wantInstalled: true}),
			ginkgo.Entry(nil, testData{dir: "005", wantID: "1.23rc1"}),
			ginkgo.Entry(nil, testData{dir: "006", wantID: "1.23rc1"}),
			ginkgo.Entry(nil, testData{dir: "007", wantID: "1.23rc1"}),
		)
	},
)
