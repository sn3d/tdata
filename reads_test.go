package tdata

import (
	"fmt"
)

func ExampleReadStr() {
	Init(nil)

	helloworld := ReadStr("folder/subfolder/helloworld.txt")
	fmt.Println(helloworld)
	// Output: Testdata
}

func ExampleReadYAML() {
	Init(nil)

	book := struct {
		Title string `yaml:"title"`
		Pages int    `yaml:"pages"`
	}{}

	ReadYAML("folder/subfolder/book.yaml", &book)

	fmt.Println(book.Title)
	fmt.Println(book.Pages)
	// Output:
	// The Mythical Man-Month
	// 272
}

func ExampleReadJSON() {
	Init(nil)

	book := struct {
		Title string `json:"title"`
		Pages int    `json:"pages"`
	}{}

	ReadJSON("folder/subfolder/book.json", &book)

	fmt.Println(book.Title)
	fmt.Println(book.Pages)
	// Output:
	// The Mythical Man-Month
	// 272
}
