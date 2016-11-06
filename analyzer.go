package deppkg

import (
	"go/build"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// ProfileName is coverprofile file name.
const ProfileName = "coverprofile"

// Analyzer analyzes coverprofile of go test
// and detects dependencies of go files and packages.
type Analyzer struct {
	// ImportPath is root package used for finding coverprofile files.
	ImportPath string
	// FileSet is set of go files and test files.
	FileSet *FileSet
	// Context is build.Context.
	Context *build.Context
}

// Analyze returns packages which are given files dependent packages
// by parsing from coverprofile files of go test.
func (a *Analyzer) Analyze() ([]*build.Package, error) {

	ctx := a.Context
	if ctx == nil {
		ctx = &build.Default
	}

	rootPkg, err := ctx.Import(a.ImportPath, "", build.FindOnly|build.IgnoreVendor)
	if err != nil {
		const msg = "cannot import root pacakge %s"
		return nil, errors.Wrapf(err, msg, a.ImportPath)
	}

	an := &analyzer{
		pkgs:        map[string]*build.Package{},
		fset:        a.FileSet,
		rootPkg:     rootPkg,
		ctx:         ctx,
		profilename: ProfileName,
	}

	if err := an.analyze(); err != nil {
		return nil, err
	}

	return an.packages(), nil
}

type analyzer struct {
	pkgs        map[string]*build.Package
	fset        *FileSet
	profiles    []*CoverProfile
	rootPkg     *build.Package
	profilename string
	ctx         *build.Context
}

func (a *analyzer) analyze() error {
	if err := a.fromTest(); err != nil {
		return err
	}

	if err := a.findProfile(); err != nil {
		return err
	}

	if err := a.fromProfile(); err != nil {
		return err
	}

	return nil
}

func (a *analyzer) fromTest() error {
	for _, tf := range a.fset.TestFiles {
		path, err := filepath.Abs(tf)
		if err != nil {
			const msg = "cannot get absolute path from %s"
			return errors.Wrapf(err, msg, tf)
		}
		dir := filepath.Dir(path)

		pkg, err := a.ctx.ImportDir(dir, build.FindOnly|build.IgnoreVendor)
		if err != nil {
			const msg = "cannot get import %s"
			return errors.Wrapf(err, msg, dir)
		}

		a.pkgs[pkg.ImportPath] = pkg
	}

	return nil
}

func (a *analyzer) findProfile() error {
	return filepath.Walk(a.rootPkg.SrcRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			const msg = "an error is occurred with finding profile"
			return errors.Wrap(err, msg)
		}

		if !info.IsDir() && info.Name() == a.profilename {

			importPath, err := ImportPath(a.rootPkg, filepath.Dir(path))
			if err != nil {
				return err
			}

			r, err := os.Open(path)
			if err != nil {
				const msg = "cannot open coverprofile file %s"
				return errors.Wrapf(err, msg, path)
			}

			p, err := ParseProfile(r, importPath)
			r.Close()

			if err != nil {
				return err
			}
			a.profiles = append(a.profiles, p)
		}

		return nil
	})
}

func (a *analyzer) fromProfile() error {

	for _, gofile := range a.fset.GoFiles {
		for _, profile := range a.profiles {

			path, err := ImportPath(a.rootPkg, gofile)
			if err != nil {
				return err
			}

			if !profile.IsDepended(path) {
				continue
			}

			if _, ok := a.pkgs[profile.ImportPath]; ok {
				continue
			}

			pkg, err := a.ctx.Import(profile.ImportPath, "", build.FindOnly|build.IgnoreVendor)
			if err != nil {
				const msg = "cannot import package %s which is got from coverprofile"
				return errors.Wrapf(err, msg, profile.ImportPath)
			}

			a.pkgs[pkg.ImportPath] = pkg
		}
	}

	return nil
}

func (a *analyzer) packages() []*build.Package {
	pkgs := make([]*build.Package, 0, len(a.pkgs))
	for _, p := range a.pkgs {
		pkgs = append(pkgs, p)
	}
	return pkgs
}
