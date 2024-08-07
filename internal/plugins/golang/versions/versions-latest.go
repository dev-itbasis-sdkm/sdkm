package versions

import (
	"context"
	"log"

	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/pkg/errors"
)

func (receiver *versions) LatestVersion(ctx context.Context) sdkmSDKVersion.SDKVersion {
	receiver.parseVersions(ctx, sdkmSDKVersion.TypeStable, receiver.reStableGroupVersions, true)

	v := receiver.cache.Load(ctx, sdkmSDKVersion.TypeStable)
	if len(v) == 0 {
		log.Fatalln(errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, "latest"))
	}

	return v[0]
}
