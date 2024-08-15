package filestorage

import (
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
)

type model struct {
	Updated  updated
	Versions map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion
}
