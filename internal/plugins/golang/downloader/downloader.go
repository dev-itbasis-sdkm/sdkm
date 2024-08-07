package downloader

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	pluginGoConsts "github.com/dev.itbasis.sdkm/internal/plugins/golang/consts"
	"github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/pkg/errors"
	"golift.io/xtractr"
)

type Downloader struct {
	httpClient *http.Client

	urlReleases string

	os   string
	arch string

	sdkDir string
}

func NewDownloader(os, arch, urlReleases, sdkDir string) *Downloader {
	return &Downloader{
		os:          os,
		arch:        arch,
		urlReleases: urlReleases,
		sdkDir:      sdkDir,
		httpClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				slog.Debug(fmt.Sprintf("'%s' redirect to '%s'...", via[0].URL, req.URL))

				if len(via) >= 10 { //nolint:mnd // TODO
					return errors.New("too many redirects")
				}

				return nil
			},
		},
	}
}

func (receiver *Downloader) Download(version string) (string, error) {
	url := receiver.URLForDownload(version)
	outFilePath := filepath.Join(receiver.sdkDir, ".download", filepath.Base(url))

	if err := os.MkdirAll(filepath.Dir(outFilePath), sdkmOs.DefaultDirMode); err != nil {
		return "", errors.Wrapf(plugin.ErrDownloadFailed, "fail make directories: %s", err.Error())
	}

	outFile, errOutFile := os.Create(outFilePath)
	if errOutFile != nil {
		return "", errors.Wrapf(plugin.ErrDownloadFailed, "fail create output file: %s", errOutFile.Error())
	}

	defer func(outFile *os.File) {
		if err := outFile.Close(); err != nil {
			panic(err)
		}
	}(outFile)

	slog.Info(fmt.Sprintf("downloading '%s' to '%s'", url, outFilePath))

	//nolint:noctx // TODO
	resp, errResp := receiver.httpClient.Get(url)
	if errResp != nil {
		return "", errors.Wrap(plugin.ErrDownloadFailed, errResp.Error())
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrapf(plugin.ErrDownloadFailed, "status code %d", resp.StatusCode)
	}

	if _, err := io.Copy(outFile, resp.Body); err != nil {
		return "", errors.Wrapf(plugin.ErrDownloadFailed, "failed copy file: %s", err.Error())
	}

	slog.Info(fmt.Sprintf("downloaded '%s' to '%s'", url, outFilePath))

	return outFilePath, nil
}

func (receiver *Downloader) Unpack(archiveFilePath, targetDir string) error {
	slog.Debug(fmt.Sprintf("unpacking '%s' to '%s'", archiveFilePath, targetDir))

	tmpDirPath, errTmpDirPath := os.MkdirTemp("", "sdkm-"+pluginGoConsts.PluginName)
	if errTmpDirPath != nil {
		return errors.Wrapf(plugin.ErrDownloadFailed, "fail create temporary dir: %s", errTmpDirPath)
	}

	defer func(path string) {
		if err := os.RemoveAll(path); err != nil {
			panic(err)
		}
	}(tmpDirPath)

	if _, _, err := xtractr.ExtractTarGzip(
		&xtractr.XFile{FilePath: archiveFilePath, OutputDir: tmpDirPath, DirMode: xtractr.DefaultDirMode, FileMode: xtractr.DefaultFileMode},
	); err != nil {
		return errors.Wrapf(plugin.ErrDownloadFailed, "extracting %s failed", archiveFilePath)
	}

	// issue https://github.com/golift/xtractr/issues/70
	if errRename := os.Rename(path.Join(tmpDirPath, "go"), targetDir); errRename != nil {
		return errors.Wrapf(plugin.ErrDownloadFailed, "failed rename: %s", errRename.Error())
	}

	return nil
}

func (receiver *Downloader) URLForDownload(version string) string {
	return fmt.Sprintf("%s/go%s.%s-%s.tar.gz", receiver.urlReleases, version, receiver.os, receiver.arch)
}
