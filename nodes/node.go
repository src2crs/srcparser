package nodes

import "go/ast"

// Node is the common interface for all node types.
type Node interface {
	Path() string
	Pos() int
	End() int
	Source() string
	Section(pos, end int) string
}

// BasicNode is a node type that encapsulates a FileInfo,
// and additionally exposes methods that create other nodes for the same file.
type BasicNode struct {
	FileInfo
	node ast.Node
}

// NewBasicNode creates a new BasicNode from a given path, content, and ast.
func NewBasicNode(fileinfo FileInfo, ast ast.Node) BasicNode {
	return BasicNode{fileinfo, ast}
}

// Pos returns the file's ast.File.Pos().
// This is a byte position with respect to its file set, cf. ast.Pos.
func (node BasicNode) Pos() int {
	return int(node.node.Pos())
}

// End returns the file's ast.File.Pos().
// This is a byte position with respect to its file set, cf. ast.Pos.
func (node BasicNode) End() int {
	return int(node.node.End())
}

// Source returns this node's source code as a string.
func (node BasicNode) Source() string {
	return node.Section(node.Pos()-node.FileInfo.Pos(), node.End()-node.FileInfo.Pos())
}

// NewFunction returns a Function node for the same file as this BasicNode.
// Expects an ast.FuncDecl that is intended to stem from this node's ast.File
func (node BasicNode) Function(f *ast.FuncDecl) Function {
	return Function{NewBasicNode(node.FileInfo, f)}
}

// NewFunction returns a Function node for the same file as this BasicNode.
// Expects an ast.FuncDecl that is intended to stem from this node's ast.File
func (node BasicNode) Import(i *ast.ImportSpec) Import {
	return Import{NewBasicNode(node.FileInfo, i)}
}
