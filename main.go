package deppkg

import "go/build"

// Main returns packages which are given files dependent packages
// by parsing from coverprofile files of go test.
// files includes *.go files and *_test.go files.
// If other files are included this function ignore them.
// pkg is an import path of root package used for finding coverprofile files.
// This function assumes dir is under $GOPATH/src directory.
// Also see Analyzer.Analyze.
func Main(ctx *build.Context, pkg string, files []string) ([]*build.Package, error) {
	return (&Analyzer{
		ImportPath: pkg,
		FileSet:    NewFileSet(files),
		Context:    ctx,
	}).Analyze()
}
