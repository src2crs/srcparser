package sourcefile

import "fmt"

// Example_print_complete_source creates a SourceFile object from a file
// and prints its complete source code.
func Example_print_complete_source() {
	// Create the SourceFile object using a file path.
	sourcefile := New("testdata/greet.go")

	// Print the file's complete source.
	fmt.Println(sourcefile.Source())

	// Output:
	// package testdata
	//
	// import (
	// 	"fmt"
	// )
	//
	// func Greet() {
	// 	fmt.Println("Hello SrcParser!")
	// }
}

// Example_print_section creates a SourceFile object from a file
// and prints a given section of its source.
func Example_print_section() {
	// Create the SourceFile object using a file path.
	sourcefile := New("testdata/greet.go")

	// Print only the import statements.
	// Note that the exact positions of the letters have to be given.
	fmt.Println(sourcefile.Section(18, 35))

	// Output:
	// import (
	// 	"fmt"
	// )
}
