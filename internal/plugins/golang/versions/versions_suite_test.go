package versions_test

import (
	"embed"
	"net/http"
	"testing"

	"github.com/itbasis/go-test-utils/v4/files"
	"github.com/itbasis/go-test-utils/v4/ginkgo"
	"github.com/onsi/gomega/ghttp"
)

//go:embed testdata/*
var testHTMLContents embed.FS

func TestVersionSuite(t *testing.T) {
	ginkgo.InitGinkgoSuite(t, "Golang Versions Suite")
}

func initFakeServer(testResponseFile string) *ghttp.Server {
	var server = ghttp.NewServer()

	server.AppendHandlers(
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/"),
			ghttp.RespondWith(http.StatusOK, files.ReadFile(testHTMLContents.ReadFile, testResponseFile)),
		),
	)

	return server
}
