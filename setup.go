package tdata

import (
	"io"
	"os"
	"path"
	"testing"
)

// global variable of temporary directory used for testing purpose.
// I made it global for easy access everywhere in tests. Because it's
// used only for testing I have no problem accept it as global
var globalTestdata *Testdata

// function do a copy of your './testdata' folder to temporary
// directory. This function ensure idempotency of tests
//
// Deprecated: this function is deprecated and you should use Init()
func InitTestdata() {
	globalTestdata = Init(nil)
}

// copy a single file from src to dst.
func cpFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.Chmod(dst, srcinfo.Mode()); err != nil {
		return err
	}

	return nil
}

// copy a whole directory recursively
func cpDir(src string, dst string) error {
	var err error
	var entries []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if entries, err = os.ReadDir(src); err != nil {
		return err
	}

	for _, entry := range entries {
		srcfp := path.Join(src, entry.Name())
		dstfp := path.Join(dst, entry.Name())

		if entry.Type().IsDir() {
			err = cpDir(srcfp, dstfp)
			if err != nil {
				return err
			}
		} else {
			err = cpFile(srcfp, dstfp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func handleError(t *testing.T, err error) {
	if t != nil {
		t.Errorf("%s - check file in testdata folder", err.Error())
		t.FailNow()
	} else {
		panic(err)
	}
}
