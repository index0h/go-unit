package unit

import (
	"path/filepath"
	"testing"
)

func TestJoin(t *testing.T) {
	t.Error(filepath.Join("namespace", "data"))
}
