package base

import (
	"io"
	"os"
	"os/exec"
	"strings"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
)

func (receiver *basePlugin) Exec(
	overrideEnv map[string]string,
	stdIn io.Reader, stdOut, stdErr io.Writer,
	args []string,
) error {
	var (
		osEnviron = os.Environ()
		envMap    = make(map[string]string, len(osEnviron)+len(overrideEnv))
	)

	for _, v := range osEnviron {
		env := strings.SplitN(v, "=", 2) //nolint:mnd // TODO
		envMap[env[0]] = env[1]
	}

	for k, v := range overrideEnv {
		envMap[k] = v
	}

	var environ = make([]string, 0, len(envMap))
	for k, v := range envMap {
		environ = append(environ, k+"="+v)
	}

	envMap["SDKM_BACKUP_PATH"] = envMap["PATH"]
	envMap["PATH"] = sdkmOs.CleanEnvPath(envMap["PATH"])

	if err := os.Setenv("PATH", envMap["PATH"]); err != nil {
		return err
	}

	cmd := exec.Command(args[0], args[1:]...) //nolint:gosec // TODO
	cmd.Stdin = stdIn
	cmd.Stdout = stdOut
	cmd.Stderr = stdErr
	cmd.Env = environ

	return cmd.Run()
}
