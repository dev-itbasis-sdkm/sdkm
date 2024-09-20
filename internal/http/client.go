package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	DefaultTimeout = 5 * time.Second

	maxRedirect = 10
)

func NewHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			slog.Debug(fmt.Sprintf("'%s' redirect to '%s'...", via[0].URL, req.URL))

			if len(via) >= maxRedirect {
				return errors.New("too many redirects")
			}

			return nil
		},
	}
}
