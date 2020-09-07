package asm

import (
	"github.com/awalterschulze/gographviz"
	"github.com/llir/llvm/ir"
)

func DiffGraphInit(src, dst *ir.Func) *gographviz.Graph {
	graph_init, _ := gographviz.ParseString(`digraph Diff {}`)
	graph_diff := gographviz.NewGraph()
	if err := gographviz.Analyse(graph_init, graph_diff); err != nil {
		panic(err)
	}
	dst.ToDotGraph(graph_diff, "Dst")
	src.ToDotGraph(graph_diff, "Src")

	return graph_diff
}
