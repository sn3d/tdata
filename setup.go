package tdata

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// global variable of temporary directory used for testing purpose.
// I made it global for easy access everywhere in tests. Because it's
// used only for testing I have no problem accept it as global
var tempDir string

// function do a copy of your './testdata' folder to temporary
// directory
//
// This function ensure idempotency of tests
func InitTestdata() {
	var err error
	if tempDir == "" {
		if tempDir, err = os.MkdirTemp("", "testdata-*"); err != nil {
			panic(err)
		}

		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		cpDir(filepath.Join(wd, "testdata"), tempDir)
	}
}

// copy a single file from src to dst.
func cpFile(src, dst string) {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		panic(err)
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		panic(err)
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		panic(err)
	}
	if srcinfo, err = os.Stat(src); err != nil {
		panic(err)
	}

	if err = os.Chmod(dst, srcinfo.Mode()); err != nil {
		panic(err)
	}
}

// copy a whole directory recursively
func cpDir(src string, dst string) {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		panic(err)
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		panic(err)
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		panic(err)
	}

	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			cpDir(srcfp, dstfp)
		} else {
			cpFile(srcfp, dstfp)
		}
	}
}
