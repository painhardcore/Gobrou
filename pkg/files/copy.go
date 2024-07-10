package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return
	}

	return dstFile.Sync()
}

func CopyDir(src, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	srcFolder, err := os.Stat(src)
	if err != nil {
		return
	}
	if !srcFolder.IsDir() {
		return fmt.Errorf("%s is not a directory, I'm disappointed.", src)
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("%s exists. U gonna regret it. ", dst)
	}

	err = os.MkdirAll(dst, srcFolder.Mode())
	if err != nil {
		return
	}

	srcFiles, err := os.ReadDir(src)
	if err != nil {
		return
	}

	for _, f := range srcFiles {
		srcPath := filepath.Join(src, f.Name())
		dstPath := filepath.Join(dst, f.Name())
		if f.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}

		} else {
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}

	}

	return
}
