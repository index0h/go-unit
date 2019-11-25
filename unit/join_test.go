package unit

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(filepath.Join("namespace", "data"))
}
