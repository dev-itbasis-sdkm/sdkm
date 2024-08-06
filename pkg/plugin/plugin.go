package plugin

import (
	"context"
	"io"
)

type GetPluginFunc func() SDKMPlugin

//nolint:interfacebloat
//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=plugin.mock.go
type SDKMPlugin interface {
	WithVersions(versions SDKVersions) SDKMPlugin
	WithBasePlugin(basePlugin BasePlugin) SDKMPlugin
	// WithDownloader(downloader) SDKMPlugin

	ListAllVersions(ctx context.Context) []SDKVersion
	ListAllVersionsByPrefix(ctx context.Context, prefix string) []SDKVersion

	LatestVersion(ctx context.Context) SDKVersion
	LatestVersionByPrefix(ctx context.Context, prefix string) (SDKVersion, error)

	Current(ctx context.Context, baseDir string) (SDKVersion, error)

	Install(ctx context.Context, baseDir string) error
	InstallVersion(ctx context.Context, version string) error

	Env(ctx context.Context, baseDir string) (map[string]string, error)
	EnvByVersion(ctx context.Context, version string) map[string]string

	Exec(ctx context.Context, baseDir string, stdIn io.Reader, stdOut, stdErr io.Writer, args []string) error
}
