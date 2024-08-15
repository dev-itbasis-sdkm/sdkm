package sdkversion

import (
	"context"
	"fmt"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=cache-storage.mock.go
type CacheStorage interface {
	fmt.Stringer

	Valid(ctx context.Context) bool

	Load(ctx context.Context) map[VersionType][]SDKVersion
	Store(ctx context.Context, versions map[VersionType][]SDKVersion)
}
