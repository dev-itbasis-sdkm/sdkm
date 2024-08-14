package versions_test

import (
	_ "embed"
	"net/http"
	"testing"

	"github.com/itbasis/go-test-utils/v4/ginkgo"
	"github.com/onsi/gomega/ghttp"
)

//go:embed testdata/all-version.html
var testHTMLContent string

func TestVersionSuite(t *testing.T) {
	ginkgo.InitGinkgoSuite(t, "Golang Versions Suite")
}

func testServer() *ghttp.Server {
	var server = ghttp.NewServer()

	server.AppendHandlers(
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/"),
			ghttp.RespondWith(http.StatusOK, testHTMLContent),
		),
	)

	return server
}
