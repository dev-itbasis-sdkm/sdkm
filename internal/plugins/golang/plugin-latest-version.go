package golang

import (
	"context"
	"log/slog"
	"strings"

	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/pkg/errors"
)

func (receiver *goPlugin) LatestVersion(ctx context.Context) sdkmSDKVersion.SDKVersion {
	return receiver.sdkVersions.LatestVersion(ctx)
}

func (receiver *goPlugin) LatestVersionByPrefix(ctx context.Context, prefix string) (sdkmSDKVersion.SDKVersion, error) {
	slog.Debug("searching for latest version by prefix: " + prefix)

	if prefix == "" {
		return receiver.LatestVersion(ctx), nil
	}

	for _, sdkVersion := range receiver.ListAllVersions(ctx) {
		if strings.HasPrefix(sdkVersion.ID, prefix) {
			receiver.enrichSDKVersion(&sdkVersion)

			return sdkVersion, nil
		}
	}

	return sdkmSDKVersion.SDKVersion{}, errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, "version by prefix "+prefix)
}
