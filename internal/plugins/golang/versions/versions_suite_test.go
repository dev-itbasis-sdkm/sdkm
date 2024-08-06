package versions_test

import (
	_ "embed"
	"net/http"
	"testing"

	sdkmTesting "github.com/dev.itbasis.sdkm/internal/testing"
	"github.com/onsi/gomega/ghttp"
)

//go:embed testdata/all-version.html
var testHTMLContent string

func TestVersionSuite(t *testing.T) {
	sdkmTesting.InitGinkgoSuite(t, "Golang Versions Suite")
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
