package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	sdkmLog "github.com/dev.itbasis.sdkm/internal/log"
	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/itbasis/go-clock/v2"
)

const (
	cacheExpirationDuration = 24 * time.Hour
)

type fileCache struct {
	Updated  fileCacheUpdated
	Versions map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion
}

func (receiver *sdkVersions) loadFromFile() {
	receiver.storeLock.Lock()
	defer receiver.storeLock.Unlock()

	filePath := receiver.filePath

	slog.Debug(fmt.Sprintf("loading cache from file: %s", filePath))

	if filePath == "" {
		return
	}

	fileInfo, errStat := os.Stat(filePath)
	if errStat != nil && os.IsNotExist(errStat) {
		return
	} else if errStat != nil {
		slog.Error("Error accessing cache file", sdkmLog.Error(errStat))

		return
	}

	if clock.Default.Now().Sub(fileInfo.ModTime()) >= cacheExpirationDuration {
		return
	}

	bytes, errReadFile := os.ReadFile(filePath)
	if errReadFile != nil {
		log.Fatalln(errReadFile) //nolint:gocritic // TODO
	}

	var fileCacheForSave fileCache

	if errUnmarshal := json.Unmarshal(bytes, &fileCacheForSave); errUnmarshal != nil {
		slog.Error("Error unmarshalling cache file", sdkmLog.Error(errUnmarshal))

		return
	}

	// TODO receiver.sdkVersions.Clear()

	for versionType, versions := range fileCacheForSave.Versions {
		receiver.sdkVersions.Store(versionType, versions)
	}
}

func (receiver *sdkVersions) saveToFile(ctx context.Context) {
	receiver.storeLock.Lock()
	defer receiver.storeLock.Unlock()

	var (
		now = clock.FromContext(ctx).Now()

		fileCacheForSave = fileCache{
			Updated:  fileCacheUpdated(now),
			Versions: map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion{},
		}
	)

	slog.Debug("Collecting cache for saving to file")

	receiver.sdkVersions.Range(
		func(key, value any) bool {
			fileCacheForSave.Versions[key.(sdkmSDKVersion.VersionType)] = value.([]sdkmSDKVersion.SDKVersion)

			return true
		},
	)

	bytes, errMarshal := json.Marshal(fileCacheForSave)
	if errMarshal != nil {
		slog.Error("failed to marshal sdk versions to file", sdkmLog.Error(errMarshal))

		return
	}

	slog.Debug(fmt.Sprintf("Save the cache to a file: %s", receiver.filePath))

	dir := filepath.Dir(receiver.filePath)
	if err := os.MkdirAll(dir, sdkmOs.DefaultDirMode); err != nil {
		slog.Error(fmt.Sprintf("failed to create directory for sdk versions to file: %s", dir), sdkmLog.Error(err))
	}

	if err := os.WriteFile(receiver.filePath, bytes, sdkmOs.DefaultFileMode); err != nil {
		slog.Error("failed to save sdk versions to file", sdkmLog.Error(err))
	}
}
