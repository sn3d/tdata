# T-Data - testdata library
[![Go Report Card](https://goreportcard.com/badge/github.com/sn3d/testdata)](https://goreportcard.com/report/github.com/sn3d/testdata)
[![codebeat badge](https://codebeat.co/badges/7cad42bc-1ddf-4b7f-ba42-542e848cffba)](https://codebeat.co/projects/github-com-sn3d-testdata-main)
[![Go Reference](https://pkg.go.dev/badge/github.com/sn3d/tdata.svg)](https://pkg.go.dev/github.com/sn3d/tdata)

This little Go library is designed to support `testdata` manipulation in 
unit tests. It enables tests that modify files to be idempotent, meaning 
they will produce the same result even if run multiple times.

That's fine if you need to read files and you're not modify it. But I have tests,
they're mutating files. I need some way how to ensure idempotency of tests.

Another issue that commonly arises when testing with files is that you cannot 
use `go embed` directly within your tests. Fortunately, the library provides 
functions like `ReadStr()`, `ReadYAML()`, and `ReadJSON()` that you can 
utilize instead. By using these functions, you can avoid having to resort to 
any kind of 'go embed' manipulation in your testing.

This library creates copy of your `testdata` folder in your `$TEMPDIR`,
for every test run. 

## Basic usage

Let's assume we have `helloworld_test.go` and `testdata` folder, where is 
`helloworld.txt`. The unit test will load the text from file, append new text 
and save it. 


```go

import . "github.com/sn3d/tdata"

func Test_HelloWorld(t *testing.T) {
   InitTestdata()

   content, err := ioutil.ReadFile(Abs("helloworld.txt")) 
   if err != nil {
      t.FailNow()
   }

   content := fmt.Sprintf("%s hello world\n", content)

   err := ioutil.WriteFile(Abs("helloworld.txt"), []byte(content), 0644)
   if err != nil {
      t.FailNow()
   }
}
```

The above example is example of idempotent file test. First, the `InitTestdata()` will create
copy of your `testdata` folder in `$TMPDIR`. The `Abs()` function will return 
absolute path to `hellowold.txt`.


### Using Read....() functions

You can use hi-level `ReadStr()`, `ReadYAML()` or `ReadJSON()` functions for 
loading files. These functions suppose to not fail. If there is problem 
with file, functions will give you no data.

Example how to read file as string:

```go
import . "github.com/sn3d/tdata"

func Test_ReadString(t *testing.T) {
   InitTestdata()

   var text string = ReadStr("helloworld.txt")

   ...
}
```

Example how to read YAML file into structure:

```go
import . "github.com/sn3d/tdata"

type Book struct {
   Title string `yaml:"title"`
   Pages int    `yaml:"pages"`
}

func Test_ReadString(t *testing.T) {
   InitTestdata()

	book := new(Book)
	ReadYAML("folder/subfolder/book.yaml", book)

   ...
}

```
