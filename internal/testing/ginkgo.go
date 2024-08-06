package testing

import (
	"log/slog"
	"testing"

	"github.com/dusted-go/logging/prettylog"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func InitGinkgoSuite(t *testing.T, name string) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	slog.SetDefault(
		slog.New(
			prettylog.New(
				&slog.HandlerOptions{Level: slog.LevelDebug},
				prettylog.WithDestinationWriter(ginkgo.GinkgoWriter),
			),
		),
	)

	ginkgo.RunSpecs(t, name)
}
