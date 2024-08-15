package cache

import (
	"context"
	"sync"

	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"golang.org/x/exp/maps"
)

type cache struct {
	storeLock sync.Mutex

	cacheStorage sdkmSDKVersion.CacheStorage
	cache        map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion
}

func NewCache() sdkmSDKVersion.Cache {
	return &cache{
		cache: map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion{},
	}
}

func (receiver *cache) String() string {
	var result = "SDKVersionCache"

	if receiver.cacheStorage != nil {
		result = result + " (" + receiver.cacheStorage.String() + ")"
	}

	return result
}

func (receiver *cache) WithExternalStore(cacheStorage sdkmSDKVersion.CacheStorage) sdkmSDKVersion.Cache {
	receiver.cacheStorage = cacheStorage
	maps.Clear(receiver.cache)

	return receiver
}

func (receiver *cache) Valid(ctx context.Context) bool {
	if receiver.cacheStorage != nil {
		return receiver.cacheStorage.Valid(ctx)
	}

	return len(receiver.cache) > 0
}

func (receiver *cache) Load(ctx context.Context, versionType sdkmSDKVersion.VersionType) []sdkmSDKVersion.SDKVersion {
	receiver.storeLock.Lock()
	defer receiver.storeLock.Unlock()

	if cacheStorage := receiver.cacheStorage; cacheStorage != nil {
		if len(receiver.cache) == 0 || !cacheStorage.Valid(ctx) {
			receiver.cache = cacheStorage.Load(ctx)
		}
	}

	var list, ok = receiver.cache[versionType]
	if !ok {
		return []sdkmSDKVersion.SDKVersion{}
	}

	return list
}

func (receiver *cache) Store(
	ctx context.Context, versionType sdkmSDKVersion.VersionType, sdkVersions []sdkmSDKVersion.SDKVersion,
) {
	receiver.storeLock.Lock()
	defer receiver.storeLock.Unlock()

	receiver.cache[versionType] = sdkVersions

	if receiver.cacheStorage != nil {
		receiver.cacheStorage.Store(ctx, receiver.cache)
	}
}
