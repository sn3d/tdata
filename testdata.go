package tdata

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v3"
)

type Testdata struct {
	TempDir string
	T       *testing.T
}

// function do a copy of your './testdata' folder to temporary
// directory and returns you instance for test
//
// This function ensure idempotency of tests
func Init(t *testing.T) *Testdata {
	var err error

	testdata := &Testdata{
		T: t,
	}

	testdata.initTempDir()

	wd, err := os.Getwd()
	if err != nil {
		testdata.handleErr(err)
	}

	err = cpDir(filepath.Join(wd, "testdata"), testdata.TempDir)
	if err != nil {
		testdata.handleErr(err)
	}

	return testdata
}

// returns you file body as []byte. If file cannot be read or
// it doesn't exits, then this function gives you empty byte array
func (t *Testdata) Read(file string) []byte {
	fileParts := strings.Split(file, "/")
	content, err := os.ReadFile(t.Abs(fileParts...))
	if err != nil {
		t.handleErr(err)
		return []byte{}
	} else {
		return content
	}
}

// returns you file body as string. If file cannot be read or
// it doesn't exits, then this function gives you empty string
func (t *Testdata) ReadStr(file string) string {
	return string(t.Read(file))
}

// function load YAML file and fill the data into
// given out
func (t *Testdata) ReadYAML(file string, out any) {
	data := t.Read(file)
	err := yaml.Unmarshal(data, out)
	if err != nil {
		t.handleErr(err)
	}
}

// function load YAML file and fill the data into
// given out
func (t *Testdata) ReadJSON(file string, out any) {
	data := t.Read(file)
	err := json.Unmarshal(data, out)
	if err != nil {
		t.handleErr(err)
	}
}

// returns you absolute path to given path in temporary directory.
func (t *Testdata) Abs(path ...string) string {
	absPath := append([]string{t.Pwd()}, path...)
	return filepath.Join(absPath...)
}

// Print working dir. In this case it will print root of
// temp directory
func (t *Testdata) Pwd() string {
	pwd, err := filepath.Abs(t.TempDir)
	if err != nil {
		t.handleErr(err)
		return t.TempDir
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
func (t *Testdata) CompareFiles(a, b string) bool {
	var err error

	aPath := filepath.Join(t.TempDir, a)
	aFile, err := os.Open(aPath)
	if err != nil {
		t.handleErr(err)
		return false
	}
	defer aFile.Close()

	bPath := filepath.Join(t.TempDir, b)
	bFile, err := os.Open(bPath)
	if err != nil {
		t.handleErr(err)
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

func (t *Testdata) handleErr(err error) {
	if t.T != nil {
		t.T.Errorf("%s - check file in testdata folder", err.Error())
		t.T.FailNow()
	} else {
		panic(err)
	}
}

func (t *Testdata) initTempDir() {
	var err error
	if t.T != nil {
		t.TempDir = t.T.TempDir()
	} else {
		t.TempDir, err = os.MkdirTemp("", "testdata-*")
		if err != nil {
			panic(err)
		}
	}

	// copy 'testdata' to temporary directory
	wd, err := os.Getwd()
	if err != nil {
		t.handleErr(err)
	}

	err = cpDir(filepath.Join(wd, "testdata"), t.TempDir)
	if err != nil {
		t.handleErr(err)
	}
}
