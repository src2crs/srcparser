package sourcefile

import "fmt"

// Example_print_errors creates a SourceFile object from a file that contains syntax errors.
// It prints the error messages that are generated by the parser.
func Example_print_errors() {
	// Create the SourceFile object using a file path.
	sourcefile := New("testdata/syntaxerror.go")

	// Print the file's source.
	fmt.Println(sourcefile.ParserErrors())

	// Output:
	// testdata/syntaxerror.go:6:38: missing ',' before newline in argument list (and 1 more errors)
}

// Example_print_no_errors creates a SourceFile object from a file that contains no errors.
// It tries to print the error messages, resulting in a message states that there were no errors.
func Example_print_no_errors() {
	// Create the SourceFile object using a file path.
	sourcefile := New("testdata/greet.go")

	// Print the parser errors of that source file.
	fmt.Println(sourcefile.ParserErrors())

	// Output:
	// testdata/greet.go: no parser errors
}
