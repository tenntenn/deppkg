package deppkg_test

import (
	"testing"

	. "github.com/tenntenn/deppkg"
)

func TestFileSet(t *testing.T) {
	files := []string{
		"hoge.go",
		"fuga/fuga.go",
		"coverprofile",
		"README",
		"hoge_test.go",
		"fuga/fuga_test.go",
	}

	gofiles := map[string]bool{
		"hoge.go":      true,
		"fuga/fuga.go": true,
	}

	testfiles := map[string]bool{
		"hoge_test.go":      true,
		"fuga/fuga_test.go": true,
	}

	fset := NewFileSet(files)

	if len(fset.GoFiles) != len(gofiles) {
		const msg = "expected number of gofiles is %d but %d"
		t.Errorf(msg, len(gofiles), len(fset.GoFiles))
	}

	for _, f := range fset.GoFiles {
		if !gofiles[f] {
			const msg = "%s was detected as go file but not expected"
			t.Errorf(msg, f)
		}
	}
	if len(fset.TestFiles) != len(testfiles) {
		const msg = "expected number of gofiles is %d but %d"
		t.Errorf(msg, len(testfiles), len(fset.TestFiles))
	}

	for _, f := range fset.TestFiles {
		if !testfiles[f] {
			const msg = "%s was detected as test file but not expected"
			t.Errorf(msg, f)
		}
	}
}
