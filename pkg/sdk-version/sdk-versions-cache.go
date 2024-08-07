package sdkversion

import (
	"context"
	"path"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=sdk-versions-cache.mock.go
type SDKVersionsCache interface {
	WithFile(filePath string) SDKVersionsCache

	Load(ctx context.Context, versionType VersionType) []SDKVersion
	Store(ctx context.Context, versionType VersionType, sdkVersions []SDKVersion)
}

func GetCacheFilePath(pluginName string) string {
	return path.Join(sdkmOs.ExecutableDir(), ".cache", pluginName+".json")
}
