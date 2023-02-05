package nodes

import (
	"fmt"
	"strings"
)

type Import struct {
	BasicNode
}

type ImportList []Import

func (imports ImportList) Source() string {
	importStrings := []string{}
	for _, i := range imports {
		importStrings = append(importStrings, fmt.Sprintf("\t%s", i.Source()))
	}
	return fmt.Sprintf("import (\n%s\n)", strings.Join(importStrings, "\n"))
}
