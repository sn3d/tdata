package tdata_test

import (
	"fmt"

	"github.com/sn3d/tdata"
)

func ExampleReadStr() {
	tdata.Setup()

	helloworld := tdata.ReadStr("folder/subfolder/helloworld.txt")
	fmt.Println(helloworld)
	// Output: Testdata
}

func ExampleReadYAML() {
	tdata.Setup()

	book := struct {
		Title string `yaml:"title"`
		Pages int    `yaml:"pages"`
	}{}

	tdata.ReadYAML("folder/subfolder/book.yaml", &book)

	fmt.Println(book.Title)
	fmt.Println(book.Pages)
	// Output:
	// The Mythical Man-Month
	// 272
}

func ExampleReadJSON() {
	tdata.Setup()

	book := struct {
		Title string `json:"title"`
		Pages int    `json:"pages"`
	}{}

	tdata.ReadJSON("folder/subfolder/book.json", &book)

	fmt.Println(book.Title)
	fmt.Println(book.Pages)
	// Output:
	// The Mythical Man-Month
	// 272
}
