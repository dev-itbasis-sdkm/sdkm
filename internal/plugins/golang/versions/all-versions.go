package versions

import (
	"context"

	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
)

func (receiver *versions) AllVersions(ctx context.Context) []sdkmPlugin.SDKVersion {
	receiver.parseVersions(ctx, sdkmPlugin.TypeStable, receiver.reStableGroupVersions, false)
	receiver.parseVersions(ctx, sdkmPlugin.TypeUnstable, receiver.reUnstableGroupVersions, false)
	receiver.parseVersions(ctx, sdkmPlugin.TypeArchived, receiver.reArchivedGroupVersions, true)

	var sdkVersions []sdkmPlugin.SDKVersion

	for _, versionType := range []sdkmPlugin.VersionType{sdkmPlugin.TypeStable, sdkmPlugin.TypeUnstable, sdkmPlugin.TypeArchived} {
		v := receiver.cache.Load(ctx, versionType)
		if len(v) == 0 {
			continue
		}

		sdkVersions = append(sdkVersions, v...)
	}

	return sdkVersions
}
