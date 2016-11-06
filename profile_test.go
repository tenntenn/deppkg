package deppkg_test

import (
	"strings"
	"testing"

	. "github.com/tenntenn/deppkg"
)

func TestParseProfile(t *testing.T) {
	const profile = `mode: set
foo/bar/hoge.go:16.78,17.108 1 1
foo/bar/hoge.go:16.78,17.108 1 1
foo/bar/fuga.go:16.78,17.108 1 1
foo/bar/piyo.go:16.78,17.108 1 1`

	expected := map[string]bool{
		"foo/bar/hoge.go":     true,
		"foo/bar/fuga.go":     true,
		"foo/bar/piyo.go":     true,
		"foo/bar/hogehoge.go": false,
	}

	cp, err := ParseProfile(strings.NewReader(profile), "foo/bar")
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	for f, b := range expected {
		if actual := cp.IsDepended(f); actual != b {
			t.Errorf("IsDepended(%s) return value is %v but %v", f, b, actual)
		}
	}

	if cp.ImportPath != "foo/bar" {
		t.Errorf("ImportPath must be given import path")
	}
}
