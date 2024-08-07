package golang

import (
	"context"
	"strings"

	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
)

func (receiver *goPlugin) ListAllVersions(ctx context.Context) []sdkmSDKVersion.SDKVersion {
	return receiver.sdkVersions.AllVersions(ctx)
}

func (receiver *goPlugin) ListAllVersionsByPrefix(ctx context.Context, prefix string) []sdkmSDKVersion.SDKVersion {
	allVersions := receiver.sdkVersions.AllVersions(ctx)

	if prefix == "" {
		return allVersions
	}

	var versions []sdkmSDKVersion.SDKVersion

	for _, sdkVersion := range allVersions {
		if strings.HasPrefix(sdkVersion.ID, prefix) {
			receiver.enrichSDKVersion(&sdkVersion)

			versions = append(versions, sdkVersion)
		}
	}

	return versions
}
