package versions

import (
	"context"
	"log"

	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/pkg/errors"
)

func (receiver *versions) LatestVersion(ctx context.Context) sdkmPlugin.SDKVersion {
	receiver.parseVersions(ctx, sdkmPlugin.TypeStable, receiver.reStableGroupVersions, true)

	v := receiver.cache.Load(ctx, sdkmPlugin.TypeStable)
	if len(v) == 0 {
		log.Fatalln(errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, "latest"))
	}

	return v[0]
}
