package cache

import (
	"context"
	"sync"

	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
)

type sdkVersions struct {
	filePath  string
	storeLock sync.Mutex

	sdkVersions sync.Map
}

func NewCacheSDKVersions() sdkmSDKVersion.SDKVersionsCache {
	return &sdkVersions{}
}

func (receiver *sdkVersions) WithFile(filePath string) sdkmSDKVersion.SDKVersionsCache {
	receiver.filePath = filePath
	receiver.loadFromFile()

	return receiver
}

func (receiver *sdkVersions) Load(
	_ context.Context, versionType sdkmSDKVersion.VersionType,
) []sdkmSDKVersion.SDKVersion {
	value, ok := receiver.sdkVersions.Load(versionType)
	if !ok {
		return []sdkmSDKVersion.SDKVersion{}
	}

	return value.([]sdkmSDKVersion.SDKVersion)
}

func (receiver *sdkVersions) Store(
	ctx context.Context, versionType sdkmSDKVersion.VersionType, sdkVersions []sdkmSDKVersion.SDKVersion,
) {
	receiver.sdkVersions.Store(versionType, sdkVersions)

	if receiver.filePath != "" {
		receiver.saveToFile(ctx)
	}
}
