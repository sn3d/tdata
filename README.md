# T-Data
[![Go Report Card](https://goreportcard.com/badge/github.com/sn3d/testdata)](https://goreportcard.com/report/github.com/sn3d/testdata)
[![codebeat badge](https://codebeat.co/badges/7cad42bc-1ddf-4b7f-ba42-542e848cffba)](https://codebeat.co/projects/github-com-sn3d-testdata-main)
[![Go Reference](https://pkg.go.dev/badge/github.com/sn3d/testdata.svg)](https://pkg.go.dev/github.com/sn3d/testdata)

This little Go library is designed to support file and folder manipulation in 
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

## Example


### Basic usage

Let's assume we have `helloworld_test.go` and `testdata` folder, where is 
`helloworld.txt`. The unit test will load the text from file, append new text 
and save it. 


```go
func Test_HelloWorld(t *testing.T) {
   tdata.Setup()

   content, err := ioutil.ReadFile(tdata.Abs("helloworld.txt")) 
   if err != nil {
      t.FailNow()
   }

   content := fmt.Sprintf("%s hello world\n", content)

   err := ioutil.WriteFile(tdata.Abs("helloworld.txt"), []byte(content), 0644)
   if err != nil {
      t.FailNow()
   }
}
```

The above example is example of idempotent file test. First, the `Setup()` will create
copy of your `testdata` folder in `$TMPDIR`. The `Abs()` function will return 
absolute path to `hellowold.txt`.


### Using Read....() functions

You can use hi-level `ReadStr()`, `ReadYAML()` or `ReadJSON()` functions for 
loading files. These functions suppose to not fail. If there is problem 
with file, functions will give you no data.

Example how to read file as string:
```go
func Test_ReadString(t *testing.T) {
   tdata.Setup()

   var text string = tdata.ReadStr("helloworld.txt")

   ...
}
```

Example how to read YAML file into structure:
```go
type Book struct {
   Title string `yaml:"title"`
   Pages int    `yaml:"pages"`
}

func Test_ReadString(t *testing.T) {
   tdata.Setup()

	book := new(Book)
	tdata.ReadYAML("folder/subfolder/book.yaml", book)

   ...
}

```
