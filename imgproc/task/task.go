package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Tasker is an interface
type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	Files  []string
}

// BuildFileList build a list of files in a directory
func BuildFileList(srcDir string) []string {
	files := []string{}
	fmt.Println("Generating file list ...")
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".jpg") {
			return nil
		}

		files = append(files, path)
		return nil
	})
	return files
}
