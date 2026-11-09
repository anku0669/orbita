// Package workspace exposes bounded local file operations.
package workspace

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Store struct{ root string }

func New(root string) (*Store, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(abs, 0o700); err != nil {
		return nil, err
	}
	return &Store{root: abs}, nil
}

func (s *Store) Resolve(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("path is required")
	}
	clean := filepath.Clean(name)
	if filepath.IsAbs(clean) {
		return "", errors.New("absolute paths are not allowed")
	}
	joined := filepath.Join(s.root, clean)
	rel, err := filepath.Rel(s.root, joined)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", errors.New("path escapes workspace")
	}
	return joined, nil
}

func (s *Store) Read(name string) ([]byte, error) {
	path, err := s.Resolve(name)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}

func (s *Store) Root() string { return s.root }
