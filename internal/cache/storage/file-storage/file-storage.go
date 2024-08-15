package filestorage

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	sdkmLog "github.com/dev.itbasis.sdkm/internal/log"
	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/itbasis/go-clock/v2"
)

const (
	cacheExpirationDuration = 24 * time.Hour
)

var (
	emptyLoadResult = map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion{}
)

type fileStorage struct {
	lock sync.Mutex

	filePath string
}

func NewFileCacheStorage(pluginName string) sdkmSDKVersion.CacheStorage {
	return NewFileCacheStorageCustomPath(path.Join(sdkmOs.ExecutableDir(), ".cache", pluginName+".json"))
}

func NewFileCacheStorageCustomPath(filePath string) sdkmSDKVersion.CacheStorage {
	slog.Debug(fmt.Sprintf("using cache with file path: %s", filePath))

	return &fileStorage{filePath: filePath}
}

func (receiver *fileStorage) String() string {
	return "FileCacheStorage[file=" + receiver.filePath + "]"
}

func (receiver *fileStorage) Valid(ctx context.Context) bool {
	filePath := receiver.filePath

	slog.Debug(fmt.Sprintf("validating with file path: %s", filePath))

	if filePath == "" {
		slog.Debug(fmt.Sprintf("file path is empty: %s", filePath))

		return false
	}

	fileInfo, errStat := os.Stat(filePath)
	if errStat != nil && os.IsNotExist(errStat) {
		slog.Debug(fmt.Sprintf("cache file not found: %s", filePath))

		return false
	} else if errStat != nil {
		slog.Error("AttrError accessing cache file", sdkmLog.AttrError(errStat))

		return false
	}

	if clock.FromContext(ctx).Now().Sub(fileInfo.ModTime()) >= cacheExpirationDuration {
		slog.Debug(fmt.Sprintf("cache file has been expired: %s", filePath))

		return false
	}

	return true
}

func (receiver *fileStorage) Load(ctx context.Context) map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	var filePath = receiver.filePath

	slog.Debug(fmt.Sprintf("loading cache from file: %s", filePath))

	if !receiver.Valid(ctx) {
		return emptyLoadResult
	}

	var bytes, errReadFile = os.ReadFile(filePath)
	if errReadFile != nil {
		slog.Error(fmt.Sprintf("error reading cache file: %s", filePath), sdkmLog.AttrError(errReadFile))

		return emptyLoadResult
	}

	var model model

	if errUnmarshal := json.Unmarshal(bytes, &model); errUnmarshal != nil {
		slog.Error(
			"error unmarshalling cache file",
			sdkmLog.AttrError(errUnmarshal),
			sdkmLog.AttrFilePath(filePath),
		)

		return emptyLoadResult
	}

	slog.Debug(fmt.Sprintf("loaded cache from file: %s", filePath))

	return model.Versions
}

func (receiver *fileStorage) Store(ctx context.Context, versions map[sdkmSDKVersion.VersionType][]sdkmSDKVersion.SDKVersion) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	filePath := receiver.filePath

	slog.Debug(fmt.Sprintf("storing cache to file: %s", filePath))

	var bytes, errMarshal = json.Marshal(
		model{
			Updated:  updated(clock.FromContext(ctx).Now()),
			Versions: versions,
		},
	)

	if errMarshal != nil {
		slog.Error(
			"error marshalling cache file",
			sdkmLog.AttrError(errMarshal),
			sdkmLog.AttrFilePath(filePath),
		)

		return
	}

	dir := filepath.Dir(filePath)
	if errMkdir := os.MkdirAll(dir, sdkmOs.DefaultDirMode); errMkdir != nil {
		slog.Error(fmt.Sprintf("error creating cache dir: %s", dir), sdkmLog.AttrError(errMkdir))

		return
	}

	if errWriteFile := os.WriteFile(filePath, bytes, sdkmOs.DefaultFileMode); errWriteFile != nil {
		slog.Error(fmt.Sprintf("error writing cache file: %s", filePath), sdkmLog.AttrError(errWriteFile))
	}
}
