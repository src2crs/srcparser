package nodes

import "go/ast"

/* TODO: error handling
 * - add a method IsValid
 * - check validity before accessing values
 * - check content size before accessing it
 * - add tests for FileInfo
 */

// FileInfo is a container type for files containing a path, file content, and ast.File.
type FileInfo struct {
	path    string
	content string
	ast     *ast.File
}

// NewFileInfo creates a new File from a given path, content, and ast.
func NewFileInfo(path string, content []byte, ast *ast.File) FileInfo {
	return FileInfo{path, string(content), ast}
}

// Path returns the file's file system path.
func (file FileInfo) Path() string {
	return file.path
}

// Pos returns the file's ast.File.Pos().
// This is a byte position with respect to its file set, cf. ast.Pos.
func (file FileInfo) Pos() int {
	return int(file.ast.Pos())
}

// End returns the file's ast.File.End().
// This is a byte position with respect to its file set, cf. ast.Pos.
func (file FileInfo) End() int {
	return int(file.ast.End())
}

// Source returns the file's complete source code as a string.
func (file FileInfo) Source() string {
	return file.content
}

// Section returns a part of the file's source code.
// The values pos and end are byte positions relative to the file.
func (file FileInfo) Section(pos, end int) string {
	return file.content[pos:end]
}
