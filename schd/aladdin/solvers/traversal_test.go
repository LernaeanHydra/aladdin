package solvers

import (
	"../cores"
	"testing"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-4
 *
 **************************************************************************************************************/

 func TestSolver_ShortestPath(t *testing.T) {
	 example   := cores.ShortestPathExample()
	 source, _ := example.GetVertex("1")
	 sink, _   := example.GetVertex("5")
	 solver    := NewShortestPathSolver(example, source, sink, NewTraversal())
	 path      := solver.ShortestPath()

	 if path.GetCost() != 20 {
		 t.Error()
	 }
 }


