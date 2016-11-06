package deppkg

import (
	"bufio"
	"io"
	"strings"

	"github.com/pkg/errors"
)

// CoverProfile has file pathes which are dependent on
// this coverprofile's target package.
type CoverProfile struct {
	m          map[string]bool
	ImportPath string
}

// ParseProfile pareses a coverprofile file.
func ParseProfile(r io.Reader, path string) (*CoverProfile, error) {
	p := &CoverProfile{
		m: map[string]bool{},
	}

	s := bufio.NewScanner(r)

	if !s.Scan() || s.Text() != "mode: set" {
		err := errors.Errorf("cannot scan 'mode: set' from %s", path)
		return nil, err
	}

	for s.Scan() {
		f := strings.Split(s.Text(), ":")[0]
		p.m[f] = true
	}

	if err := s.Err(); err != nil {
		const msg = "cannot scan coverprofile file %s"
		return nil, errors.Wrapf(err, msg, path)
	}

	p.ImportPath = path

	return p, nil
}

// IsDepended returns whether given file is dependent on this package.
func (p *CoverProfile) IsDepended(file string) bool {
	return p.m[file]
}
