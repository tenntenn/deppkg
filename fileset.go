package deppkg

import (
	"path/filepath"
	"strings"
)

// FileSet is set of go files and testfiles.
type FileSet struct {
	GoFiles   []string
	TestFiles []string
}

// NewFileSet create a new FileSet object from given files.
// It ignores files which are not go file and Go's test file.
func NewFileSet(files []string) *FileSet {
	var fset FileSet

	for _, f := range files {
		if filepath.Ext(f) != ".go" {
			continue
		}

		if strings.HasSuffix(f, "_test.go") {
			fset.TestFiles = append(fset.TestFiles, f)
		} else {
			fset.GoFiles = append(fset.GoFiles, f)
		}
	}

	return &fset
}
