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
	"time"

	"github.com/dev.itbasis.sdkm/internal/cache"
	sdkmLog "github.com/dev.itbasis.sdkm/internal/log"
	pluginGoConsts "github.com/dev.itbasis.sdkm/internal/plugins/golang/consts"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
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

	cache sdkmPlugin.SDKVersionsCache
}

func NewVersions(urlReleases string) sdkmPlugin.SDKVersions {
	return &versions{
		urlReleases: urlReleases,
		cache: cache.NewCacheSDKVersions().
			WithFile(sdkmPlugin.GetCacheFilePath(pluginGoConsts.PluginName)),

		httpClient: &http.Client{
			Timeout: 5 * time.Second, //nolint:mnd //
		},
		reStableGroupVersions:   regexp.MustCompile(`<h2 id="stable">.*?<h2 id=`),
		reUnstableGroupVersions: regexp.MustCompile(`<h2 id="unstable">.*?<div.*?id="archive"`),
		reArchivedGroupVersions: regexp.MustCompile(`id="archive">.+?</article`),
		reGoVersion:             regexp.MustCompile(`id="go(.+?)"`),
	}
}

func (receiver *versions) WithCache(cache sdkmPlugin.SDKVersionsCache) sdkmPlugin.SDKVersions {
	receiver.cache = cache

	return receiver
}

func (receiver *versions) getContent(url string) (string, error) {
	slog.Debug(fmt.Sprintf("getting content for url: %s", url))

	resp, err := receiver.httpClient.Get(url) //nolint:noctx
	if err != nil {
		return "", err //nolint:wrapcheck //
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("Error closing body after receiving content", sdkmLog.Error(err))
		}
	}()

	body, err := io.ReadAll(resp.Body)

	return strings.ReplaceAll(string(body), "\n", ""), err
}

func (receiver *versions) parseVersions(
	ctx context.Context,
	versionType sdkmPlugin.VersionType,
	reGroupVersions *regexp.Regexp,
	cleanContent bool,
) {
	if len(receiver.cache.Load(ctx, versionType)) > 0 {
		return
	}

	receiver.muParsing.Lock()
	defer receiver.muParsing.Unlock()

	if receiver.contentReleases == "" {
		receiver.contentReleases, _ = receiver.getContent(receiver.urlReleases)

		if cleanContent {
			defer func() {
				receiver.contentReleases = ""
			}()
		}
	}

	content := reGroupVersions.FindString(receiver.contentReleases)
	if content == "" {
		return
	}

	submatch := receiver.reGoVersion.FindAllStringSubmatch(content, -1)

	var sdkVersions = make([]sdkmPlugin.SDKVersion, len(submatch))

	for i, row := range submatch {
		if row[1] != "" {
			sdkVersion := sdkmPlugin.SDKVersion{ID: row[1], Type: versionType}

			sdkVersions[i] = sdkVersion
		}
	}

	receiver.cache.Store(ctx, versionType, sdkVersions)
}
