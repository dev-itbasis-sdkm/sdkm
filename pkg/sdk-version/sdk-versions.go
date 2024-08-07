package sdkversion

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=sdk-versions.mock.go
type SDKVersions interface {
	WithCache(cache SDKVersionsCache) SDKVersions

	AllVersions(ctx context.Context) []SDKVersion
	LatestVersion(ctx context.Context) SDKVersion
}
