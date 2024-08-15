package filestorage_test

import (
	filestorage "github.com/dev.itbasis.sdkm/internal/cache/storage/file-storage"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.DescribeTable(
	"NewFileCacheStorage", func(pluginName, wantString string) {
		gomega.Expect(filestorage.NewFileCacheStorage(pluginName).String()).To(
			gomega.SatisfyAll(
				gomega.HavePrefix("FileCacheStorage[file="),
				gomega.HaveSuffix(wantString),
			),
		)
	},
	ginkgo.Entry(nil, "go", "/.cache/go.json]"),
	ginkgo.Entry(nil, "jvm", "/.cache/jvm.json]"),
	ginkgo.Entry(nil, "jvm-openjdk", "/.cache/jvm-openjdk.json]"),
)

var _ = ginkgo.DescribeTable(
	"NewFileCacheStorageCustomPath", func(pluginName, wantString string) {
		gomega.Expect(filestorage.NewFileCacheStorageCustomPath(pluginName).String()).To(
			gomega.SatisfyAll(
				gomega.HavePrefix("FileCacheStorage[file="),
				gomega.HaveSuffix(wantString),
			),
		)
	},
	ginkgo.Entry(nil, "go", "go]"),
	ginkgo.Entry(nil, "/.cache/go.json", "/.cache/go.json]"),
)
