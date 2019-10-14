package dsl

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/dsl/component"
	"strconv"
	"strings"
)

func GenerateDigraph(wires [][]string, symbolsTable map[string]component.Component) string {
	var sb strings.Builder

	sb.WriteString("digraph function_expanded_graph {\n")

	for k, v := range symbolsTable {
		sb.WriteString(strconv.Quote(k) + " [label=" + strconv.Quote(fmt.Sprintf("%s", v)) + "];\n")
	}

	for _, w := range wires {
		sb.WriteString(generateWire(w) + ";\n")
	}

	sb.WriteString("}\n")
	return sb.String()
}

func generateWire(wire []string) string {
	quoted := make([]string, len(wire))
	for i, c := range wire {
		quoted[i] = strconv.Quote(c)
	}
	return strings.Join(quoted, " -> ")
}
