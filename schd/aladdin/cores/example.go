package cores

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-4
 *
 **************************************************************************************************************/

/*
 * @see <a href=
 *      "https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm">Solver's
 *      Algorithm (Wikipedia)</a> <br>*
 */
 func ShortestPathExample() *Graph {
	 graph := NewGraph()
	 v1    := NewVertex("1")
	 v2    := NewVertex("2")
	 v3    := NewVertex("3")
	 v4    := NewVertex("4")
	 v5    := NewVertex("5")
	 v6    := NewVertex("6")

	 graph.AddVertex(v1)
	 graph.AddVertex(v2)
	 graph.AddVertex(v3)
	 graph.AddVertex(v4)
	 graph.AddVertex(v5)
	 graph.AddVertex(v6)

	 e16   := NewCostEdge(14,NewIntCapacity(1), v1, v6)
	 e12   := NewCostEdge(7, NewIntCapacity(1), v1, v2)
	 e13   := NewCostEdge(9, NewIntCapacity(1), v1, v3)
	 e23   := NewCostEdge(10,NewIntCapacity(1), v2, v3)
	 e24   := NewCostEdge(15,NewIntCapacity(1), v2, v4)
	 e34   := NewCostEdge(11,NewIntCapacity(1), v3, v4)
	 e36   := NewCostEdge(2, NewIntCapacity(1), v3, v6)
	 e45   := NewCostEdge(6, NewIntCapacity(1), v4, v5)
	 e65   := NewCostEdge(9, NewIntCapacity(1), v6, v5)

	 graph.AddEdge(e12)
	 graph.AddEdge(e13)
	 graph.AddEdge(e16)
	 graph.AddEdge(e23)
	 graph.AddEdge(e24)
	 graph.AddEdge(e34)
	 graph.AddEdge(e36)
	 graph.AddEdge(e45)
	 graph.AddEdge(e65)

	 return graph
 }

/*
* @see <a href=
*      "https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm>Solver's
*      Algorithm (Wikipedia)</a> <br>*
*/
func MaxFlowExample() *Graph {
	graph := NewGraph()
	a := NewVertex("A")
	b := NewVertex("B")
	c := NewVertex("C")
	d := NewVertex("D")
	e := NewVertex("E")
	f := NewVertex("F")
	g := NewVertex("G")

	graph.AddVertex(a)
	graph.AddVertex(b)
	graph.AddVertex(c)
	graph.AddVertex(d)
	graph.AddVertex(e)
	graph.AddVertex(f)
	graph.AddVertex(g)


	ab := NewCostEdge(3,NewIntCapacity(1), a, b)
	ad := NewCostEdge(3,NewIntCapacity(1), a, d)
	bc := NewCostEdge(4,NewIntCapacity(1), b, c)
	ca := NewCostEdge(3,NewIntCapacity(1), c, a)
	cd := NewCostEdge(1,NewIntCapacity(1), c, d)
	ce := NewCostEdge(2,NewIntCapacity(1), c, e)
	de := NewCostEdge(2,NewIntCapacity(1), d, e)
	df := NewCostEdge(6,NewIntCapacity(1), d, f)
	eb := NewCostEdge(1,NewIntCapacity(1), e, b)
	eg := NewCostEdge(1,NewIntCapacity(1), e, g)
	fg := NewCostEdge(9,NewIntCapacity(1), f, g)

	graph.AddEdge(ab)
	graph.AddEdge(bc)
	graph.AddEdge(ad)
	graph.AddEdge(ca)
	graph.AddEdge(cd)
	graph.AddEdge(ce)
	graph.AddEdge(de)
	graph.AddEdge(df)
	graph.AddEdge(eb)
	graph.AddEdge(eg)
	graph.AddEdge(fg)

	return graph
}
