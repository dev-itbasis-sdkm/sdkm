package versions

import (
	"context"

	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
)

func (receiver *versions) AllVersions(ctx context.Context) []sdkmSDKVersion.SDKVersion {
	if !receiver.cache.Valid(ctx) {
		receiver.parseVersions(ctx, sdkmSDKVersion.TypeStable, receiver.reStableGroupVersions, false)
		receiver.parseVersions(ctx, sdkmSDKVersion.TypeUnstable, receiver.reUnstableGroupVersions, false)
		receiver.parseVersions(ctx, sdkmSDKVersion.TypeArchived, receiver.reArchivedGroupVersions, true)
	}

	var sdkVersions []sdkmSDKVersion.SDKVersion

	for _, versionType := range []sdkmSDKVersion.VersionType{sdkmSDKVersion.TypeStable, sdkmSDKVersion.TypeUnstable, sdkmSDKVersion.TypeArchived} {
		v := receiver.cache.Load(ctx, versionType)
		if len(v) == 0 {
			continue
		}

		sdkVersions = append(sdkVersions, v...)
	}

	return sdkVersions
}
