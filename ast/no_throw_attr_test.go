package ast

import (
	"testing"
)

func TestNoThrowAttr(t *testing.T) {
	nodes := map[string]Node{
		`0x7fa1488273a0 <line:7:4, line:11:4>`: &NoThrowAttr{
			Address:  "0x7fa1488273a0",
			Position: "line:7:4, line:11:4",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
