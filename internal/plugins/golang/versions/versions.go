package versions

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
	"sync"

	sdkmCache "github.com/dev.itbasis.sdkm/internal/cache"
	sdkmHttp "github.com/dev.itbasis.sdkm/internal/http"
	sdkmLog "github.com/dev.itbasis.sdkm/internal/log"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
)

type versions struct {
	urlReleases string

	httpClient *http.Client

	muParsing sync.Mutex

	contentReleases string

	reStableGroupVersions   *regexp.Regexp
	reUnstableGroupVersions *regexp.Regexp
	reArchivedGroupVersions *regexp.Regexp
	reGoVersion             *regexp.Regexp

	cache sdkmSDKVersion.SDKVersionsCache
}

func NewVersions(urlReleases string) sdkmSDKVersion.SDKVersions {
	return &versions{
		urlReleases: urlReleases,
		cache:       sdkmCache.NewCacheSDKVersions(),

		httpClient: sdkmHttp.NewHTTPClient(),

		reStableGroupVersions:   regexp.MustCompile(`<h2 id="stable">.*?<h2`),
		reUnstableGroupVersions: regexp.MustCompile(`<h2 id="unstable">.*?<div.*?id="archive"`),
		reArchivedGroupVersions: regexp.MustCompile(`id="archive">.+?</article`),
		reGoVersion:             regexp.MustCompile(`id="go(.+?)"`),
	}
}

func (receiver *versions) WithCache(cache sdkmSDKVersion.SDKVersionsCache) sdkmSDKVersion.SDKVersions {
	slog.Debug(fmt.Sprintf("setting cache: %s", cache))

	receiver.cache = cache

	return receiver
}

func (receiver *versions) getContent(url string) (string, error) {
	slog.Debug(fmt.Sprintf("getting content for url: %s", url))

	resp, errGet := receiver.httpClient.Get(url) //nolint:noctx // TODO
	if errGet != nil {
		return "", errGet //nolint:wrapcheck // TODO
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("Error closing body after receiving content", sdkmLog.Error(err))
		}
	}()

	body, errReadAll := io.ReadAll(resp.Body)

	content := strings.ReplaceAll(string(body), "\n", "")

	slog.Debug(fmt.Sprintf("received content size: %d", len(content)))

	return content, errReadAll
}

func (receiver *versions) parseVersions(
	ctx context.Context,
	versionType sdkmSDKVersion.VersionType,
	reGroupVersions *regexp.Regexp,
	cleanContent bool,
) {
	if len(receiver.cache.Load(ctx, versionType)) > 0 {
		slog.Debug(fmt.Sprintf("cache hit for version: %s", versionType))

		return
	}

	slog.Debug(fmt.Sprintf("cache miss for version: %s", versionType))

	receiver.muParsing.Lock()
	defer receiver.muParsing.Unlock()

	if receiver.contentReleases == "" {
		receiver.contentReleases, _ = receiver.getContent(receiver.urlReleases)

		// if slog.Default().Enabled(ctx, slog.LevelDebug) {
		// if err := os.WriteFile(".url-releases.html", []byte(receiver.contentReleases), sdkmOs.DefaultFileMode); err != nil {
		// 	slog.Error("Error creating .url-releases.html", sdkmLog.Error(err))
		// }
		// }

		if cleanContent {
			defer func() {
				receiver.contentReleases = ""
			}()
		}
	}

	var content = reGroupVersions.FindString(receiver.contentReleases)
	if content == "" {
		slog.Debug(fmt.Sprintf("content is empty for version: %s", versionType))

		return
	}

	slog.Debug(fmt.Sprintf("found groups for version type: %s", versionType))

	var (
		submatch    = receiver.reGoVersion.FindAllStringSubmatch(content, -1)
		sdkVersions = make([]sdkmSDKVersion.SDKVersion, len(submatch))
	)

	for i, row := range submatch {
		if row[1] != "" {
			sdkVersion := sdkmSDKVersion.SDKVersion{ID: row[1], Type: versionType}

			sdkVersions[i] = sdkVersion
		}
	}

	slog.Debug(fmt.Sprintf("found %d SDK versions for version type: %s", len(sdkVersions), versionType))

	receiver.cache.Store(ctx, versionType, sdkVersions)
}
