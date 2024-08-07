package sdkversion_test

import (
	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"Print", func() {
		ginkgo.DescribeTable(
			"Strange models, but need to check", func(model sdkmSDKVersion.SDKVersion, expected string) {
				gomega.Expect(model.Print()).To(gomega.Equal(expected))
			},
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1"}, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeStable}, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeUnstable}, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeArchived}, ""),
		)

		ginkgo.DescribeTable(
			"correct models", func(model sdkmSDKVersion.SDKVersion, expected string) {
				gomega.Expect(model.Print()).To(gomega.Equal(expected))
			},
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable}, "1"),
			ginkgo.Entry(
				nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable, Installed: true}, "1 [installed]",
			),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable}, "1 (unstable)"),
			ginkgo.Entry(
				nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable, Installed: true},
				"1 (unstable) [installed]",
			),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeArchived}, "1 (archived)"),
			ginkgo.Entry(
				nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeArchived, Installed: true},
				"1 (archived) [installed]",
			),
		)
	},
)

var _ = ginkgo.Describe(
	"Print with options", func() {
		ginkgo.DescribeTable(
			"Strange models, but need to check",
			func(model sdkmSDKVersion.SDKVersion, outType, outInstalled, outNotInstalled bool, expected string) {
				gomega.Expect(model.PrintWithOptions(outType, outInstalled, outNotInstalled)).To(gomega.Equal(expected))
			},
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, false, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, false, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, false, true, true, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, true, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, true, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, true, true, true, ""),

			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1"}, false, false, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1"}, false, true, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1"}, false, true, true, "1 [not installed]"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1"}, true, false, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1"}, true, true, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1"}, true, true, true, "1 [not installed]"),

			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeStable}, false, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeStable}, false, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeStable}, false, true, true, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeStable}, true, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeStable}, true, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeStable}, true, true, true, ""),

			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeUnstable}, false, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeUnstable}, false, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeUnstable}, false, true, true, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeUnstable}, true, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeUnstable}, true, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeUnstable}, true, true, true, ""),

			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeArchived}, false, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeArchived}, false, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeArchived}, false, true, true, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeArchived}, true, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeArchived}, true, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{Type: sdkmSDKVersion.TypeArchived}, true, true, true, ""),
		)

		ginkgo.DescribeTable(
			"correct",
			func(model sdkmSDKVersion.SDKVersion, outType, outInstalled, outNotInstalled bool, expected string) {
				gomega.Expect(model.PrintWithOptions(outType, outInstalled, outNotInstalled)).To(gomega.Equal(expected))
			},
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, false, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, false, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, false, true, true, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, true, false, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, true, true, false, ""),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{}, true, true, true, ""),

			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable}, false, false, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable}, false, true, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable}, false, true, true, "1 [not installed]"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable}, true, false, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable}, true, true, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeStable}, true, true, true, "1 [not installed]"),

			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable}, false, false, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable}, false, true, false, "1"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable}, false, true, true, "1 [not installed]"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable}, true, false, false, "1 (unstable)"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable}, true, true, false, "1 (unstable)"),
			ginkgo.Entry(nil, sdkmSDKVersion.SDKVersion{ID: "1", Type: sdkmSDKVersion.TypeUnstable}, true, true, true, "1 (unstable) [not installed]"),
		)
	},
)
