package golang

import (
	"context"
	"fmt"
	"os"
	"path"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	pluginGoConsts "github.com/dev.itbasis.sdkm/internal/plugins/golang/consts"
)

func (receiver *goPlugin) Env(ctx context.Context, baseDir string) (map[string]string, error) {
	sdkVersion, errCurrent := receiver.Current(ctx, baseDir)
	if errCurrent != nil {
		return map[string]string{}, errCurrent
	}

	return receiver.EnvByVersion(ctx, sdkVersion.ID), nil
}

func (receiver *goPlugin) EnvByVersion(_ context.Context, version string) map[string]string {
	goCacheDir := path.Join(sdkmOs.UserHomeDir(), pluginGoConsts.PluginName, version)
	goBin := path.Join(goCacheDir, "bin")

	originPath := os.Getenv("SDKM_PATH_ORIGIN")
	if originPath == "" {
		originPath = os.Getenv("PATH")
	}

	envs := map[string]string{
		"SDKM_PATH_ORIGIN":   os.Getenv("PATH"),
		"SDKM_GOROOT_ORIGIN": os.Getenv("GOROOT"),
		"SDKM_GOPATH_ORIGIN": os.Getenv("GOPATH"),
		"SDKM_GOBIN_ORIGIN":  os.Getenv("GOBIN"),
		//
		"GOROOT": receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginName, version),
		"GOPATH": goCacheDir,
		"GOBIN":  goBin,
		"PATH": fmt.Sprintf(
			"%s%c%s%c%s%c%s",
			sdkmOs.ExecutableDir(),
			os.PathListSeparator,
			path.Join(receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginName, version), "bin"),
			os.PathListSeparator,
			goBin,
			os.PathListSeparator,
			originPath,
		),
	}

	return envs
}
