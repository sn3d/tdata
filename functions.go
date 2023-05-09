package tdata

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// returns you absolute path to given path in temporary directory.
func Abs(path ...string) string {
	absPath := append([]string{Pwd()}, path...)
	return filepath.Join(absPath...)
}

// Print working dir. In this case it will print root of
// temp directory
func Pwd() string {
	pwd, err := filepath.Abs(tempDir)
	if err != nil {
		return tempDir
	} else {
		return pwd
	}
}

// compare content o files, function returns true if they're
// matching, otherwise it returns false.
//
// the a and b are relative paths to need to temporary testdata
// directory
//
// The false is returned also when any error occurs.
//
// The function  isn't optimal but it serves only for
// testing purposes
func CompareFiles(a, b string) bool {
	var err error

	aPath := filepath.Join(tempDir, a)
	aFile, err := os.Open(aPath)
	if err != nil {
		return false
	}
	defer aFile.Close()

	bPath := filepath.Join(tempDir, b)
	bFile, err := os.Open(bPath)
	if err != nil {
		return false
	}
	defer bFile.Close()

	aData := make([]byte, 1024)
	bData := make([]byte, 1024)

	for {
		aSize, err1 := aFile.Read(aData)
		bSize, err2 := bFile.Read(bData)

		if err1 == io.EOF && err2 == io.EOF {
			break
		}

		if err1 != nil || err2 != nil {
			return false
		}

		if aSize != bSize {
			return false
		}

		if !bytes.Equal(aData, bData) {
			return false
		}
	}

	return true
}
