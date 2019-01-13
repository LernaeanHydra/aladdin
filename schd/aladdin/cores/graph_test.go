package cores

import "testing"

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-4
 *
 **************************************************************************************************************/

 func TestNewGraph(t *testing.T) {
	 graph := ShortestPathExample()

	 _, vv := graph.GetVertex("1")

	 if vv != nil {
	 	t.Error();
	 }

	 _, nvv := graph.GetVertex("7")

	 if nvv == nil {
	 	t.Error()
	 }

	 _, ve := graph.GetEdge("1-2")

	 if ve != nil {
		 t.Error();
	 }

	 _, nve := graph.GetEdge("1-7")

	 if nve == nil {
		 t.Error()
	 }
 }