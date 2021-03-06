package ast

import (
	"testing"
)

func TestContinueStmt(t *testing.T) {
	nodes := map[string]Node{
		`0x1e044e0 <col:20>`: &ContinueStmt{
			Address:  "0x1e044e0",
			Position: "col:20",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
