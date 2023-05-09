package tdata

import (
	"fmt"
)

func ExampleCompareFiles() {
	InitTestdata()

	fmt.Println(CompareFiles("compare/f1.txt", "compare/f2.txt"))
	fmt.Println(CompareFiles("compare/f1.txt", "compare/f3.txt"))

	// Output:
	// true
	// false
}
