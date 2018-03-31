package main

import (
	"os"
	"fmt"
	"archive/zip"
	"path/filepath"
	"strings"
	"io"
	"time"
)

func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

func backupDS2DS3()  {
	sources := [2]string {`\DarkSoulsIII`, `\DarkSoulsII`}
	fmt.Println(len(sources))
	for i:=0;i < len(sources);i++{
		target := os.Getenv("AppData") + `\ds_backups`+sources[i]+time.Now().Format(".01.02.2006.15-04-05")+`.zip`
		fmt.Println(target)
		zipit(os.Getenv("AppData")+sources[i], target)
	}
}

func main() {
	os.Mkdir(os.Getenv("AppData")+`\ds_backups`, 0700)
	backupDS2DS3()
}
