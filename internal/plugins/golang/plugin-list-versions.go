package golang

import (
	"context"
	"strings"

	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
)

func (receiver *goPlugin) ListAllVersions(ctx context.Context) []sdkmPlugin.SDKVersion {
	return receiver.sdkVersions.AllVersions(ctx)
}

func (receiver *goPlugin) ListAllVersionsByPrefix(ctx context.Context, prefix string) []sdkmPlugin.SDKVersion {
	allVersions := receiver.sdkVersions.AllVersions(ctx)

	if prefix == "" {
		return allVersions
	}

	var versions []sdkmPlugin.SDKVersion

	for _, sdkVersion := range allVersions {
		if strings.HasPrefix(sdkVersion.ID, prefix) {
			receiver.enrichSDKVersion(&sdkVersion)

			versions = append(versions, sdkVersion)
		}
	}

	return versions
}
