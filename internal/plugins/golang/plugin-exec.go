package golang

import (
	"context"
	"io"

	"github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/pkg/errors"
)

func (receiver *goPlugin) Exec(
	ctx context.Context, baseDir string, stdIn io.Reader, stdOut, stdErr io.Writer, args []string,
) error {
	environ, errEnviron := receiver.Env(ctx, baseDir)
	if errEnviron != nil {
		return errors.Wrapf(plugin.ErrExecuteFailed, "failed get environment: %s", errEnviron.Error())
	}

	if errExec := receiver.basePlugin.Exec(environ, stdIn, stdOut, stdErr, args); errExec != nil {
		return errors.Wrapf(plugin.ErrExecuteFailed, "failed exec: %s", errExec.Error())
	}

	return errEnviron
}
