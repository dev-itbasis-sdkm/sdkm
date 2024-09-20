package golang

import (
	"context"
	"fmt"
	"log/slog"

	pluginGoConsts "github.com/dev.itbasis.sdkm/internal/plugins/golang/consts"
	"github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/pkg/errors"
)

func (receiver *goPlugin) Install(ctx context.Context, baseDir string) error {
	sdkVersion, errCurrent := receiver.Current(ctx, baseDir)
	if errCurrent != nil {
		return errors.Wrapf(plugin.ErrSDKInstall, "failed get current version: %s", errCurrent.Error())
	}

	if sdkVersion.Installed {
		slog.Info(fmt.Sprintf("SDK already installed: %s", sdkVersion.ID))

		return nil
	}

	return receiver.InstallVersion(ctx, sdkVersion.ID)
}

func (receiver *goPlugin) InstallVersion(_ context.Context, version string) error {
	if receiver.basePlugin.HasInstalled(pluginGoConsts.PluginName, version) {
		slog.Info("SDK already installed: " + version)

		return nil
	}

	archiveFilePath, errDownload := receiver.downloader.Download(version)
	if errDownload != nil {
		return errors.Wrap(plugin.ErrSDKInstall, errDownload.Error())
	}

	if errUnpack := receiver.downloader.Unpack(
		archiveFilePath, receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginName, version),
	); errUnpack != nil {
		return errors.Wrap(plugin.ErrSDKInstall, errUnpack.Error())
	}

	return nil
}
