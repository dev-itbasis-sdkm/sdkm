package cache

import (
	"context"
	"sync"

	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
)

type sdkVersions struct {
	filePath  string
	storeLock sync.Mutex

	sdkVersions sync.Map
}

func NewCacheSDKVersions() sdkmPlugin.SDKVersionsCache {
	return &sdkVersions{}
}

func (receiver *sdkVersions) WithFile(filePath string) sdkmPlugin.SDKVersionsCache {
	receiver.filePath = filePath
	receiver.loadFromFile()

	return receiver
}

func (receiver *sdkVersions) Load(_ context.Context, versionType sdkmPlugin.VersionType) []sdkmPlugin.SDKVersion {
	value, ok := receiver.sdkVersions.Load(versionType)
	if !ok {
		return []sdkmPlugin.SDKVersion{}
	}

	return value.([]sdkmPlugin.SDKVersion)
}

func (receiver *sdkVersions) Store(
	ctx context.Context, versionType sdkmPlugin.VersionType, sdkVersions []sdkmPlugin.SDKVersion,
) {
	receiver.sdkVersions.Store(versionType, sdkVersions)

	if receiver.filePath != "" {
		receiver.saveToFile(ctx)
	}
}
