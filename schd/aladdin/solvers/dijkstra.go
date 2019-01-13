package solvers

import (
	"k8s.io/kubernetes/schd/aladdin/cores"
	"container/list"
)

/************************************************************************************************************
 * <p>
 * Worst case: O(|E| + |V| log |V|)
 * <p>
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-4
 * @note      please catch exception by yourself
 *
 *
 **************************************************************************************************************/

 type Dijkstra struct {
	unvisited *list.List
 }

 func NewDijkstra() *Dijkstra {
	unvisited := list.New().Init()
	return &Dijkstra{
		unvisited: unvisited,
	}
 }

 func (d *Dijkstra) Push(v cores.Vertex) {
 	// list just initial
 	if d.unvisited.Len() == 0 {
		d.unvisited.PushBack(v)
		return
	}

	newValue := v.GetDistance()

	for e := d.unvisited.Front(); e != nil; e = e.Next() {
		thisValue := e.Value.(cores.Vertex)
		if newValue < thisValue.GetDistance() {
			d.unvisited.InsertBefore(v, e)
			break
		}
	}

	 lastElem := d.unvisited.Back().Value.(cores.Vertex)
	 if newValue > lastElem.GetDistance() {
		d.unvisited.PushBack(v)
	}
 }

 func (d *Dijkstra) Next(b bool, m map[string]string, v cores.Vertex) {
	if b == false {
		d.Push(v)
	}
 }

 func (d *Dijkstra) Len() int {
	return d.unvisited.Len()
 }

 func (d *Dijkstra) Pop() cores.Vertex{
	vertex := d.unvisited.Front().Value.(cores.Vertex)
	d.unvisited.Remove(d.unvisited.Front())
	return vertex
 }
