package golang_test

import (
	"context"
	"os"
	"path"
	"strings"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmPluginGo "github.com/dev.itbasis.sdkm/internal/plugins/golang"
	pluginGoConsts "github.com/dev.itbasis.sdkm/internal/plugins/golang/consts"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = ginkgo.Describe(
	"EnvByVersion", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.BeforeEach(
			func() {
				ginkgo.GinkgoT().Setenv("SDKM_PATH_ORIGIN", "")
			},
		)

		ginkgo.DescribeTable(
			"success", func(sdkVersion, wantSDKPath, wantGoCachePath string) {
				var (
					originPath      = os.Getenv("PATH")
					originPaths     = strings.Split(originPath, string(os.PathListSeparator))
					countOriginPath = len(originPaths)

					mockController = gomock.NewController(ginkgo.GinkgoT())
					mockBasePlugin = sdkmPlugin.NewMockBasePlugin(mockController)

					sdkVersionDir = path.Join("sdk", sdkVersion)
				)

				mockBasePlugin.EXPECT().GetSDKDir().Return(sdkVersionDir).AnyTimes()
				mockBasePlugin.EXPECT().GetSDKVersionDir(pluginGoConsts.PluginName, sdkVersion).Return(sdkVersionDir).AnyTimes()

				var envs = sdkmPluginGo.GetPlugin().
					WithBasePlugin(mockBasePlugin).
					EnvByVersion(context.Background(), sdkVersion)

				gomega.Expect(envs).To(
					gomega.SatisfyAll(
						gomega.HaveLen(8),

						gomega.HaveKeyWithValue("SDKM_PATH_ORIGIN", originPath),
					),
				)

				var actualPaths = strings.Split(envs["PATH"], string(os.PathListSeparator))

				gomega.Expect(originPaths).To(gomega.HaveLen(countOriginPath))
				gomega.Expect(actualPaths).To(gomega.HaveLen(countOriginPath + 3))
				gomega.Expect(actualPaths[0]).To(gomega.Equal(wantSDKPath))
				gomega.Expect(actualPaths[1]).To(gomega.Equal(wantGoCachePath))
				gomega.Expect(actualPaths[2]).To(gomega.Equal(sdkmOs.ExecutableDir()))
			},
			ginkgo.Entry(nil, "1.23.0", path.Join("sdk", "1.23.0", "bin"), path.Join(sdkmOs.UserHomeDir(), "go", "1.23.0", "bin")),
			ginkgo.Entry(nil, "1.20.1", path.Join("sdk", "1.20.1", "bin"), path.Join(sdkmOs.UserHomeDir(), "go", "1.20.1", "bin")),
		)
	},
)
