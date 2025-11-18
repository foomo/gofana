package gofana

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

func List(roots []string) ([]string, error) {
	var paths []string

	for _, root := range roots {
		if err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
			if err != nil || !d.IsDir() {
				return err
			}
			if strings.HasPrefix(d.Name(), ".") {
				return fs.SkipDir
			}
			if d.Name() == "gofana" {
				paths = append(paths, path.Join(p, "main.go"))
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	return paths, nil
}
