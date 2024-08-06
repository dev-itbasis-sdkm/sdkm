package golang

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/pkg/errors"
)

func (receiver *goPlugin) LatestVersion(ctx context.Context) sdkmPlugin.SDKVersion {
	return receiver.sdkVersions.LatestVersion(ctx)
}

func (receiver *goPlugin) LatestVersionByPrefix(ctx context.Context, prefix string) (sdkmPlugin.SDKVersion, error) {
	slog.Debug(fmt.Sprintf("searching for latest version by prefix: %s", prefix))

	if prefix == "" {
		return receiver.LatestVersion(ctx), nil
	}

	for _, sdkVersion := range receiver.ListAllVersions(ctx) {
		if strings.HasPrefix(sdkVersion.ID, prefix) {
			receiver.enrichSDKVersion(&sdkVersion)

			return sdkVersion, nil
		}
	}

	return sdkmPlugin.SDKVersion{}, errors.Wrap(
		sdkmPlugin.ErrSDKVersionNotFound, fmt.Sprintf("version by prefix %s", prefix),
	)
}
