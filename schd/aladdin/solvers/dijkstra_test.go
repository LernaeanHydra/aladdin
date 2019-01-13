package solvers

import (
	"../cores"
	"testing"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-16
 *
 **************************************************************************************************************/

 func TestSolver_ShortestPath2(t *testing.T) {
	 example   := cores.ShortestPathExample()
	 source, _ := example.GetVertex("1")
	 sink, _   := example.GetVertex("5")
	 solver    := NewShortestPathSolver(example, source, sink, NewDijkstra())
	 path      := solver.ShortestPath()

	 if path.GetCost() != 20 {
		 t.Error()
	 }
 }
