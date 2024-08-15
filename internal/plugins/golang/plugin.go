package golang

import (
	"runtime"

	sdkmCache "github.com/dev.itbasis.sdkm/internal/cache"
	file_storage "github.com/dev.itbasis.sdkm/internal/cache/storage/file-storage"
	pluginBase "github.com/dev.itbasis.sdkm/internal/plugins/base"
	pluginGoConsts "github.com/dev.itbasis.sdkm/internal/plugins/golang/consts"
	pluginsGoDownloader "github.com/dev.itbasis.sdkm/internal/plugins/golang/downloader"
	pluginGoVersions "github.com/dev.itbasis.sdkm/internal/plugins/golang/versions"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
)

type goPlugin struct {
	sdkmPlugin.SDKMPlugin

	basePlugin  sdkmPlugin.BasePlugin
	sdkVersions sdkmSDKVersion.SDKVersions
	downloader  *pluginsGoDownloader.Downloader
}

func GetPlugin() sdkmPlugin.SDKMPlugin {
	basePlugin := pluginBase.NewBasePlugin()
	downloader := pluginsGoDownloader.NewDownloader(
		runtime.GOOS, runtime.GOARCH, pluginGoConsts.URLReleases, basePlugin.GetSDKDir(),
	)
	cache := sdkmCache.NewCache().
		WithExternalStore(file_storage.NewFileCacheStorage(pluginGoConsts.PluginName))

	sdkVersions := pluginGoVersions.NewVersions(pluginGoConsts.URLReleases).
		WithCache(cache)

	return &goPlugin{
		basePlugin:  basePlugin,
		downloader:  downloader,
		sdkVersions: sdkVersions,
	}
}

func (receiver *goPlugin) WithBasePlugin(basePlugin sdkmPlugin.BasePlugin) sdkmPlugin.SDKMPlugin {
	receiver.basePlugin = basePlugin
	receiver.downloader = pluginsGoDownloader.NewDownloader(
		runtime.GOOS, runtime.GOARCH, pluginGoConsts.URLReleases, basePlugin.GetSDKDir(),
	)

	return receiver
}

func (receiver *goPlugin) WithVersions(versions sdkmSDKVersion.SDKVersions) sdkmPlugin.SDKMPlugin {
	receiver.sdkVersions = versions

	return receiver
}

func (receiver *goPlugin) enrichSDKVersion(sdkVersion *sdkmSDKVersion.SDKVersion) {
	if sdkVersion == nil {
		return
	}

	sdkVersion.Installed = sdkVersion.Installed ||
		receiver.basePlugin.HasInstalled(pluginGoConsts.PluginName, sdkVersion.ID)
}
