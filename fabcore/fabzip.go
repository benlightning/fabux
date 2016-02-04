// Copyright 2014 fabux authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package fabcore

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Files2Zip(dir, name string) {
	zipName := fmt.Sprintf("%s.zip", name)
	currentDir, _ := os.Getwd()
	file, _ := os.Create(zipName)
	defer file.Close()
	err := os.Chdir(filepath.Dir(dir))
	if err != nil {
		os.Exit(1)
	}
	dir = filepath.Base(dir)
	zipbuf := zip.NewWriter(file)
	defer zipbuf.Close()

	walk := func(dir string, info os.FileInfo, err error) error {
		if err != nil {
			error2curDir(currentDir, zipName, file)
			return err
		}
		if info.IsDir() {
			return nil
		}

		src, _ := os.Open(dir)
		defer src.Close()
		h := &zip.FileHeader{Name: dir, Method: zip.Deflate, Flags: 0x800}
		filename, _ := zipbuf.CreateHeader(h)
		io.Copy(filename, src)
		zipbuf.Flush()
		return nil
	}
	filepath.Walk(dir, walk)
	os.Chdir(currentDir)
}

func error2curDir(currentDir, zipName string, file *os.File) {
	os.Chdir(currentDir)
	file.Close()
	os.Remove(zipName)
}
