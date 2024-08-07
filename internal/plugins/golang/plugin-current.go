package golang

import (
	"context"

	"github.com/dev.itbasis.sdkm/internal/plugins/golang/modfile"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
)

func (receiver *goPlugin) Current(ctx context.Context, baseDir string) (sdkmSDKVersion.SDKVersion, error) {
	goModFile, errGoModFile := modfile.ReadGoModFile(baseDir)
	if errGoModFile != nil {
		return sdkmSDKVersion.SDKVersion{}, errGoModFile //nolint:wrapcheck // TODO
	}

	var (
		sdkVersion sdkmSDKVersion.SDKVersion
		err        error
	)

	if toolchain := goModFile.Toolchain; toolchain != nil {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, toolchain.Name[2:])
	} else {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, goModFile.Go.Version)
	}

	if err != nil {
		return sdkmSDKVersion.SDKVersion{}, err //nolint:wrapcheck // TODO
	}

	receiver.enrichSDKVersion(&sdkVersion)

	return sdkVersion, nil
}
