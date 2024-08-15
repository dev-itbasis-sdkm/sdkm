package sdkversion

import (
	"context"
	"fmt"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=cache.mock.go
type Cache interface {
	fmt.Stringer

	WithExternalStore(cacheStorage CacheStorage) Cache

	Valid(ctx context.Context) bool
	Load(ctx context.Context, versionType VersionType) []SDKVersion
	Store(ctx context.Context, versionType VersionType, sdkVersions []SDKVersion)
}
