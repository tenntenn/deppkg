package deppkg

import (
	"go/build"
	"path/filepath"

	"github.com/pkg/errors"
)

func ImportPath(rootPkg *build.Package, f string) (string, error) {
	af, err := filepath.Abs(f)
	if err != nil {
		const msg = "cannot get absolute path from %s"
		return "", errors.Wrapf(err, msg, f)
	}

	path, err := filepath.Rel(rootPkg.SrcRoot, af)
	if err != nil {
		const msg = "cannot get relative path %s and %s"
		return "", errors.Wrapf(err, msg, rootPkg.SrcRoot, af)
	}

	return path, nil
}
