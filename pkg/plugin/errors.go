package plugin

import "errors"

var (
	ErrSDKVersionNotFound = errors.New("SDK version not found")

	ErrSDKInstall     = errors.New("SDK install error")
	ErrDownloadFailed = errors.New("download failed")

	ErrExecuteFailed = errors.New("execute failed")
)
