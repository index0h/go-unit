package unit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestJoin(t *testing.T) {
	rootPath, err := ioutil.TempDir("", "")

	if err != nil {
		t.Error(err)
	}

	nsPath := filepath.Join(rootPath, "namespace1")

	if err := os.Mkdir(nsPath, 0777); err != nil {
		t.Error(err)
	}

	if file, err := os.Create("file.go"); err != nil {
		t.Error(err)
	} else {
		t.Error(file)
	}

	t.Errorf("%+v", findAllFolders(rootPath))
}

func findAllFolders(path string) []string {
	result := []string{}

	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				result = append(result, path)
			}

			return nil
		},
	)

	if err != nil {
		panic(err)
	}

	sort.Strings(result)

	return result
}
