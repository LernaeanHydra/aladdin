/**
 *  Copyright (2018, ) Institute of Software, Chinese Academy of Sciences
 */

package main

import (
	"k8s.io/kubernetes/schd/aladdin/cores"
	"k8s.io/kubernetes/schd/aladdin/solvers"
	"fmt"
)

func main() {
	//testSolver()
	//testAntiAffinity()
	example   := cores.MaxFlowExample()
	source, _ := example.GetVertex("A")
	sink, _   := example.GetVertex("G")
	solver    := solvers.NewSMaxFlowSolver(example, *source, *sink, solvers.NewDijkstra())
	//fmt.Println(solver.ShortestPath().GoString())
	fmt.Println(solver.MaxFlow().GoString())
}


func testAntiAffinity() {
	aa := cores.NewAntiAffinity()
	values, err := aa.ReadLine("/home/alibaba-data/prod/app_interference")
	if err != nil {
		fmt.Println(err)
	}
	for k, _ := range values {
		fmt.Println(k)
	}
}

func testSolver () {
	example   := cores.ShortestPathExample()
	source, _ := example.GetVertex("1")
	sink, _   := example.GetVertex("5")
	//solver    := solvers.NewShortestPathSolver(example, source, sink, solvers.NewTraversal())
	solver    := solvers.NewShortestPathSolver(example, *source, *sink, solvers.NewDijkstra())
	path      := solver.ShortestPath()

	fmt.Println(path.GoString())
}