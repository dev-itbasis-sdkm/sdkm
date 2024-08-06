package modfile

import (
	"os"
	"path"
	"path/filepath"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	"golang.org/x/mod/modfile"
)

func ReadGoModFile(baseDir string) (modfile.File, error) {
	var (
		goModFilePath             = path.Join(baseDir, "go.mod")
		goModFileRelPath, errPath = filepath.Rel(sdkmOs.Pwd(), goModFilePath)
	)

	if errPath != nil {
		return modfile.File{}, errPath
	}

	bytes, errRead := os.ReadFile(goModFilePath)
	if errRead != nil {
		return modfile.File{}, errRead
	}

	file, errParse := modfile.Parse(goModFileRelPath, bytes, nil)
	if errParse != nil {
		return modfile.File{}, errParse
	}

	return *file, nil
}
