package golang

import (
	"context"

	"github.com/dev.itbasis.sdkm/internal/plugins/golang/modfile"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
)

func (receiver *goPlugin) Current(ctx context.Context, baseDir string) (sdkmPlugin.SDKVersion, error) {
	goModFile, errGoModFile := modfile.ReadGoModFile(baseDir)
	if errGoModFile != nil {
		return sdkmPlugin.SDKVersion{}, errGoModFile //nolint:wrapcheck
	}

	var (
		sdkVersion sdkmPlugin.SDKVersion
		err        error
	)

	if toolchain := goModFile.Toolchain; toolchain != nil {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, toolchain.Name[2:])
	} else {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, goModFile.Go.Version)
	}

	if err != nil {
		return sdkmPlugin.SDKVersion{}, err //nolint:wrapcheck
	}

	receiver.enrichSDKVersion(&sdkVersion)

	return sdkVersion, nil
}
