package cores

import (
	"container/list"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-3
 *
 **************************************************************************************************************/

type Path struct {
	cost  int
	edges *list.List
}

func NewPath() *Path {
	return &Path{cost: 0,
		edges: list.New()}
}

func NewInvalidPath() *Path {
	return &Path{cost: -1,
		edges: list.New()}
}

func (path *Path) GetCost() int {
	return path.cost
}

// we would not check parameters
func (path *Path) AddEdge(edge Edge) {
	path.cost += edge.cost
 	//path.edges[edge.name] = *edge
 	path.edges.PushBack(edge)
}

func (path *Path) GetEdges() *list.List {
	return path.edges
}