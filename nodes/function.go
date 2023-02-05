package nodes

import "go/ast"

// Function is a nodetype for functions in a source file.
type Function struct {
	BasicNode
}

// Name returns the function's name as a string.
func (function Function) Name() string {
	return function.node.(*ast.FuncDecl).Name.Name
}
