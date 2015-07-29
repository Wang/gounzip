package unzip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	Silence bool
)

func Do(zipfile, targetDir string) error {
	if targetDir == "" {
		targetDir = "./" //默认在本目录
	}

	if targetDir != "./" {
		if !IsDir(targetDir) {
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				return errors.New("target directory not exists .")
			}
		}
	}

	if zipfile == "" {
		return errors.New("Usage : unzip test.zip")
	}

	reader, err := zip.OpenReader(zipfile)
	if err != nil {
		return errors.New(fmt.Sprintf("Error, Read zip file: %s", err.Error()))
		os.Exit(1)
	}
	defer reader.Close()

	Printf("Archive: %s", zipfile)

	for _, f := range reader.Reader.File {
		zipped, err := f.Open()
		if err != nil {
			return errors.New(fmt.Sprintf("Error, Open zipped file: %s", err.Error()))
		}
		defer zipped.Close()

		path := filepath.Join(targetDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			Printf("Creating directory: %s", path)
		} else {
			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
			if err != nil {
				return errors.New(fmt.Sprintf("Error, Create file: %s", err.Error()))
			}
			defer writer.Close()

			if _, err = io.Copy(writer, zipped); err != nil {
				return errors.New(fmt.Sprintf("Error, Write file content: %s", err.Error()))
			}
			Printf("Decompressing: %s", path)
		}
	}

	return nil
}

func Printf(format string, v ...interface{}) {
	if Silence {
		return
	}
	fmt.Printf(format+"\n", v...)
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return false
	}
	return fi.IsDir()
}
