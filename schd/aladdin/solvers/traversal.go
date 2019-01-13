package solvers

import (
	"k8s.io/kubernetes/schd/aladdin/cores"
	"container/list"
)

/************************************************************************************************************
 * <p>
 * Worst case: O(|E| + |V|^2)
 * <p>
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-4
 * @note      only for acyclic graph
 *            please catch exception by yourself
 *
 **************************************************************************************************************/

 type Traversal struct {
 	unvisited *list.List
 }

 func NewTraversal() *Traversal {
 	unvisited := list.New().Init()
 	return &Traversal{
 		unvisited: unvisited,
	}
 }

 func (t *Traversal) Push(v cores.Vertex) {
 	t.unvisited.PushBack(v)
 }

 func (t *Traversal) Next(b bool, m map[string]string, v cores.Vertex) {
	t.unvisited.PushBack(v)
 }

 func (t *Traversal) Len() int {
	return t.unvisited.Len()
 }

 func (t *Traversal) Pop() cores.Vertex{
	 vertex := t.unvisited.Front().Value.(cores.Vertex)
	 t.unvisited.Remove(t.unvisited.Front())
	 return vertex
 }

