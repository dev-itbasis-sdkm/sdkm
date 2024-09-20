package golang

import (
	"context"
	"fmt"
	"log/slog"
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
	var (
		goCacheDir = receiver.getGoCacheDir(version)
		goBin      = path.Join(goCacheDir, "bin")
		originPath = os.Getenv("SDKM_PATH_ORIGIN")
	)

	slog.Debug("originPath: " + originPath)

	if originPath == "" {
		originPath = os.Getenv("PATH")

		slog.Debug("originPath (using env.PATH): " + originPath)
	}

	var envs = map[string]string{
		"SDKM_PATH_ORIGIN":   originPath,
		"SDKM_GOROOT_ORIGIN": os.Getenv("GOROOT"),
		"SDKM_GOPATH_ORIGIN": os.Getenv("GOPATH"),
		"SDKM_GOBIN_ORIGIN":  os.Getenv("GOBIN"),
		//
		"GOROOT": receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginName, version),
		"GOPATH": goCacheDir,
		"GOBIN":  goBin,
		"PATH": sdkmOs.AddBeforePath(
			originPath,
			path.Join(receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginName, version), "bin"),
			goBin,
			sdkmOs.ExecutableDir(),
		),
	}

	slog.Debug(fmt.Sprintf("envs: %v", envs))

	return envs
}
