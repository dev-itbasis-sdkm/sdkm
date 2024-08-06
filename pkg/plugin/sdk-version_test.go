package plugin_test

import (
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"ToString()", func() {

		ginkgo.DescribeTable(
			"Strange models, but need to check", func(model sdkmPlugin.SDKVersion, expected string) {
				gomega.Expect(model.String()).To(gomega.Equal(expected))
			},
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{}, ""),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{ID: "1"}, "1"),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{Type: sdkmPlugin.TypeStable}, ""),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{Type: sdkmPlugin.TypeUnstable}, ""),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{Type: sdkmPlugin.TypeArchived}, ""),
		)

		ginkgo.DescribeTable(
			"Correct models", func(model sdkmPlugin.SDKVersion, expected string) {
				gomega.Expect(model.String()).To(gomega.Equal(expected))
			},
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{}, ""),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{ID: "1", Type: sdkmPlugin.TypeStable}, "1"),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{ID: "1", Type: sdkmPlugin.TypeStable, Installed: true}, "1 [installed]"),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{ID: "1", Type: sdkmPlugin.TypeUnstable}, "1 (unstable)"),
			ginkgo.Entry(
				nil, sdkmPlugin.SDKVersion{ID: "1", Type: sdkmPlugin.TypeUnstable, Installed: true},
				"1 (unstable) [installed]",
			),
			ginkgo.Entry(nil, sdkmPlugin.SDKVersion{ID: "1", Type: sdkmPlugin.TypeArchived}, "1 (archived)"),
			ginkgo.Entry(
				nil, sdkmPlugin.SDKVersion{ID: "1", Type: sdkmPlugin.TypeArchived, Installed: true},
				"1 (archived) [installed]",
			),
		)
	},
)
