package cores

import (
	"errors"
	"log"
	"k8s.io/api/core/v1"
	"fmt"
	"math/rand"
	"strconv"
	//"github.com/onsi/gomega/matchers/support/goraph/edge"
	//"MCMF/data"
)

/************************************************************************************************************
 * A graph G = (V, E) consists of a set of vertices, V, and a set of edges, E. Each edge is a pair
 * (v, w), where v, w ∈ V. Edges are sometimes referred to as arcs. If the pair is ordered, then the
 * graph is directed. Directed graphs are sometimes referred to as digraphs. Vertex w is adjacent to
 * v if and only if (v, w) ∈ E. In an undirected graph with edge (v, w), and hence (w, v), w is adjacent
 * to v and v is adjacent to w. Sometimes an edge has a third component, known as either a weight or
 * a cost.
 *
 * A path in a graph is a sequence of vertices w1, w2, w3, ... , wN such that (wi, wi+1) ∈ E
 * for 1 ≤ i < N. The length of such a path is the number of edges on the path, which is
 * equal to N − 1. We allow a path from a vertex to itself; if this path contains no edges, then
 * the path length is 0. This is a convenient way to define an otherwise special case. If the
 * graph contains an edge (v, v) from a vertex to itself, then the path v, v is sometimes referred
 * to as a loop. The graphs we will consider will generally be loopless. A simple path is a
 * path such that all vertices are distinct, except that the first and last could be the same
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-2
 *
 **************************************************************************************************************/

/***********************************
 *
 *    Graph cores
 *
 ***********************************/

type Graph struct {
	vertices map[string]Vertex
	edges    map[string]Edge
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]Vertex),
		edges:    make(map[string]Edge),
	}
}

//todo how to construct a greph with pods and nodes
func (graph *Graph) InitGraphVertex(nodes []*v1.Node,  pods map[string]*v1.Pod) error{
	if nodes == nil || len(nodes) == 0 || pods == nil || len(pods) == 0 {
		fmt.Printf("pods or nodes con't be empty when constructing map")
		return errors.New("pods or nodes con't be empty when constructing map")
	}
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	source := NewVertex("source")
	graph.AddVertex(source)
	sink := NewVertex("sink")
	graph.AddVertex(sink)

	for _, pod := range pods {
		cpu := 0
		for _, container := range pod.Spec.Containers {
			cpu = int(container.Resources.Requests.Cpu().Value())
		}
		vertex := NewVertex(pod.Name)
		graph.AddVertex(vertex)
		var edge *Edge
		edge = NewCostEdge(r1.Intn(10), &IntCapacity{cpu}, source, vertex)
		// edge := NewCostEdge(r1.Intn(10), &IntCapacity{r1.Intn(10)}, source, vertex)
		graph.AddEdge(edge)
	}

	for _, node := range nodes {
		cpu := int(node.Status.Allocatable.Cpu().Value())

		vertex := NewVertex(node.Name)
		graph.AddVertex(vertex)
		for _, podVtx := range source.GetOutEdges(){
			edge := NewCostEdge(r1.Intn(10), NewIntCapacityWithCapacity(podVtx.GetCapacity()), podVtx.to, vertex)
			graph.AddEdge(edge)
		}
		edge := NewCostEdge(r1.Intn(10), &IntCapacity{cpu}, vertex, sink)
		graph.AddEdge(edge)

	}
	return nil
}

func (graph *Graph) UpdateGraghForDijkstra(path Path) Capacity{


	minFlow := NewIntCapacity(-1) // current min cost of path

	// we need to update capacity for graph
	// here get min capacity of the path
	for e := path.GetEdges().Front(); e != nil; e = e.Next() {
		edge := e.Value.(Edge)
		if minFlow.IsNull() || edge.GetCapacity().Less(minFlow) {
			minFlow = NewIntCapacityWithCapacity(edge.GetCapacity())
		}

	}
	// set new capacity for every edge of this path, if edge cost is 0, then delete it from vertex's outEdge
	for e := path.GetEdges().Front(); e != nil; e = e.Next() {
		// remove flow capacity from the edge
		value := e.Value.(Edge)
		// get the edge in gragh, we need to update its capacity
		edge,_ := graph.GetEdge(GetEdgeName(
			value.GetFrom(), value.GetTo()))
		if err := edge.GetCapacity().Sub(minFlow); err != nil{
			fmt.Printf("capacity can't below 0 when updating capacity")
			panic("capacity can't below 0 when updating capacity")
		}

		// if the edge's capacity is zero, delete this edge from prior vertex's outEdge
		//if edge.GetCost() == 0 {
		//	name := cores.GetEdgeName(value.GetFrom(), value.GetTo())
		//	delete(edge.GetFrom().GetOutEdges(), name)
		//	degs[name] = name
		//}

		// todo why here get reverse name
		//name := cores.GetEdgeName(value.GetTo(), value.GetFrom())


		// todo why??
		//if _, ok := degs[name]; ok {
		//	break
		//}

		//edge, err := s.graph.GetEdge(name)
		//if err != nil {
		//	newEdge := cores.NewCostEdge(min, value.GetTo(), value.GetFrom())
		//	newEdge.SetCost(min)
		//	s.graph.AddEdge(newEdge)
		//} else {
		//	edge.AddCost(edge.GetCost() + min)  // todo ???
		//}
	}
	//s.graph.PrintGragh()
	return minFlow
}

func (graph *Graph) UpdateGraghForMaxFlow(path Path) Capacity{

	fmt.Println("Need to update graph after getting path :"+path.GoString())

	edgeNum := path.GetEdges().Len()
	incrementFlow := NewIntCapacity(-1) // current min cost of path

	//如果是直接调度成功的
	index := 0;
	if edgeNum == 3 {
		for e := path.GetEdges().Front(); e != nil; e = e.Next() {
			edge := e.Value.(Edge)
			if index == 0 {
				incrementFlow = NewIntCapacityWithCapacity(edge.GetCapacity())
				break
			}

		}
		for e := path.GetEdges().Front(); e != nil; e = e.Next() {
			// remove flow capacity from the edge
			value := e.Value.(Edge)
			// get the edge in gragh, we need to update its capacity
			edge, _ := graph.GetEdge(GetEdgeName(
				value.GetFrom(), value.GetTo()))
			if err := edge.GetCapacity().Sub(incrementFlow); err != nil {
				fmt.Printf("capacity can't below 0 when updating capacity")
				panic("capacity can't below 0 when updating capacity")
			}
		}

	}else if edgeNum == 4{  //反悔整个任务
		for e := path.GetEdges().Front(); e != nil; e = e.Next() {
			edge := e.Value.(Edge)
			if index == 0 {
				incrementFlow = NewIntCapacityWithCapacity(edge.GetCapacity())
			}
			if index == 2{
				incrementFlow.Sub(edge.GetMaxCapacity())
				break
			}
			index ++

		}
		index = 0
		for e := path.GetEdges().Front(); e != nil; e = e.Next() {
			// remove flow capacity from the edge
			value := e.Value.(Edge)
			// get the edge in gragh, we need to update its capacity
			edge, _ := graph.GetEdge(GetEdgeName(
				value.GetFrom(), value.GetTo()))
			if index <2 {
				if err := edge.GetCapacity().Sub(value.GetCapacity()); err != nil {
					fmt.Printf("capacity can't below 0 when updating capacity")
					panic("capacity can't below 0 when updating capacity")
				}
				// 更新反悔的机器上的容量
				if index == 1 {
					machineVertex := edge.GetTo()
					for _, machineOutEdge := range machineVertex.GetOutEdges(){
						edge, _ := graph.GetEdge(GetEdgeName(
							machineOutEdge.GetFrom(), machineOutEdge.GetTo()))
						edge.GetCapacity().Sub(incrementFlow)
					}
				}
			}
			if index >=2 && index <4{
				edge.GetCapacity().Add(edge.GetMaxCapacity())
			}

			index++
		}

	}else if edgeNum == 5{  // 反悔任务到别的机器上
		regretFlow := NewIntCapacity(0)
		flowChange := NewIntCapacity(0)
		for e := path.GetEdges().Front(); e != nil; e = e.Next() {
			edge := e.Value.(Edge)
			if index == 0 {
				incrementFlow = NewIntCapacityWithCapacity(edge.GetCapacity())
				flowChange.Add(edge.GetCapacity())
			}
			if index == 2{
				regretFlow.Add(edge.GetMaxCapacity())
				flowChange.Sub(edge.GetMaxCapacity())
				break
			}
			index ++

		}
		index = 0
		for e := path.GetEdges().Front(); e != nil; e = e.Next() {
			// remove flow capacity from the edge
			value := e.Value.(Edge)
			// get the edge in gragh, we need to update its capacity
			edge, _ := graph.GetEdge(GetEdgeName(
				value.GetFrom(), value.GetTo()))
			if index <2 {
				if err := edge.GetCapacity().Sub(incrementFlow); err != nil {
					fmt.Printf("capacity can't below 0 when updating capacity")
					panic("capacity can't below 0 when updating capacity")
				}
				// 更新反悔的机器上的容量
				if index == 1 {
					machineVertex := edge.GetTo()
					for _, machineOutEdge := range machineVertex.GetOutEdges(){
						edge, _ := graph.GetEdge(GetEdgeName(
							machineOutEdge.GetFrom(), machineOutEdge.GetTo()))
						edge.GetCapacity().Sub(flowChange)
					}
				}
			}
			if index ==2 {
				edge.GetCapacity().Add(edge.GetMaxCapacity())
			}

			if index >2 {
				edge.GetCapacity().Sub(regretFlow)
			}

			index++
		}
	}else{  // 暂不支持
		panic("invalid path type")
	}

	return incrementFlow
}

func (graph *Graph) AddVertex(vertex *Vertex) {
	if IsNullVertex(vertex) &&
		IsVertexExist(graph.GetVertices(), vertex.GetName()) {
			log.Print("Null vertex or the vertex exists")
			return
	}
	graph.vertices[vertex.name] = *vertex
}

func (graph *Graph) GetVertices() map[string]Vertex {
	return graph.vertices
}

func (graph *Graph) GetVertex(name string) (vertex *Vertex, error error) {
	if IsNullString(name) ||
		!IsVertexExist(graph.GetVertices(), name) {
			log.Print("String is null or vertex with the specified name does not exist")
			error = errors.New("no vertex with the specified name found")
	}
	value := graph.vertices[name]
	vertex = &value
	return
}

func (graph *Graph) RemoveVertex(name string) (vertex *Vertex, error error) {
	if IsNullString(name) ||
		!IsVertexExist(graph.GetVertices(), name) {
		log.Print("String is null or vertex with the specified name does not exist")
		error = errors.New("no vertex with the specified name found")
	}
	vertex, _ = graph.GetVertex(name)
	delete(graph.vertices, name)
	vertex.inEdges = make(map[string]Edge)
	vertex.outEdges = make(map[string]Edge)
	return
}

func (graph *Graph) AddEdge(edge *Edge) {
	if IsNullEdge(edge) ||
		IsNullVertex(edge.from) ||
		IsNullVertex(edge.to) {
			log.Print("Invalid edge")
			return
	}

	name := GetEdgeName(
		    edge.from, edge.to)
	graph.edges[name] = *edge
	edge.to.inEdges[name] = *edge
	edge.from.outEdges[name] = *edge
}

func (graph *Graph) GetEdges() map[string]Edge {
	return graph.edges
}

func (graph *Graph) GetEdge(name string) (edge Edge, error error) {
	if IsNullString(name) ||
		!IsEdgeExist(graph.GetEdges(), name) {
			log.Print("String is null or edge with the specified name does not exist")
			error = errors.New("no edge with the specified name found")

	}
	edge = graph.edges[name]
	return
}

func (graph *Graph) RemoveEdge(name string) (edge Edge, error error) {
	if IsNullString(name) ||
		!IsEdgeExist(graph.GetEdges(), name) {
		log.Print("String is null or edge with the specified name does not exist")
		error = errors.New("no edge with the specified name found")
	}
	edge, _ = graph.GetEdge(name)
	delete(graph.edges, name)
	delete(edge.from.outEdges, name)
	delete(edge.to.inEdges, name)
	return
}

func (graph *Graph) PrintGragh(){
	fmt.Println("**** Start print graph ****")
	for _, edge := range graph.edges {
		fmt.Println(edge.name+ ": " +edge.capacity.GoString())
	}
	fmt.Println("**** End print graph ****")
}


/************************************
 *
 *    Vertex cores
 *
 ***********************************/

 type Vertex struct {
	name     string // vertex name
	distance int
 	outEdges map[string]Edge
	inEdges  map[string]Edge
 }

 func NewVertex(name string) *Vertex {
	return &Vertex{
		name: name,
		distance: 0,
		outEdges: make(map[string]Edge),
		inEdges:  make(map[string]Edge),
	}
 }

 func (vertex *Vertex) SetDistance(distance int) {
	vertex.distance = distance
 }

 func (vertex *Vertex) GetDistance() int {
	return vertex.distance
 }

 func (vertex *Vertex) GetName() string {
	return vertex.name
 }

 func (vertex *Vertex) GetOutEdges() map[string]Edge {
	return vertex.outEdges
 }

 func (vertex *Vertex) GetInEdges() map[string]Edge {
	return vertex.inEdges
 }

 func (vertex *Vertex) GetOutEdge(name string) (edge Edge, error error) {
	 if IsNullString(name) ||
		 !IsEdgeExist(vertex.GetOutEdges(), name) {
		 log.Print("String is null or edge with the specified name does not exist")
		 error = errors.New("no edge with the specified name found")
	 }
	 edge = vertex.outEdges[name]
	 return
 }

 func (vertex *Vertex) GetInEdge(name string) (edge Edge, error error) {
	if IsNullString(name) ||
		!IsEdgeExist(vertex.GetInEdges(), name) {
		log.Print("String is null or edge with the specified name does not exist")
		error = errors.New("no edge with the specified name found")
	}
	edge = vertex.inEdges[name]
	return
 }

 // please do not add the same edge twice
 // we would not check it
 func (vertex *Vertex) AddOutEdge(edge *Edge) {
	if IsNullEdge(edge)  {
		log.Print("Null edge")
		return
	}
	vertex.outEdges[GetEdgeName(edge.from, edge.to)] = *edge
 }

// please do not add the same edge twice
// we would not check it
 func (vertex *Vertex) AddInEdge(edge *Edge) {
	if IsNullEdge(edge) {
		log.Print("Null edge")
		return
	}
	vertex.inEdges[GetEdgeName(edge.from, edge.to)] = *edge
 }

 func (vertex *Vertex) RemoveOutEdge(name string) (edge Edge, error error) {
	if IsNullString(name) ||
		!IsEdgeExist(vertex.GetOutEdges(), name) {
			log.Print("String is null or edge with the specified name does not exist")
			error = errors.New("no edge with the specified name found")
	}
	edge, _ = vertex.GetOutEdge(name)
	delete(vertex.outEdges, name)
	return
 }

 func (vertex *Vertex) RemoveInEdge(name string) (edge Edge, error error) {
	if IsNullString(name) ||
		!IsEdgeExist(vertex.GetInEdges(), name) {
			log.Print("String is null or edge with the specified name does not exist")
			error = errors.New("no edge with the specified name found")
	}
	edge, _ = vertex.GetInEdge(name)
	delete(vertex.inEdges, name)
	return
 }

/************************************
 *
 *    Edge cores
 *
 ************************************/
type Edge struct {
	name     string
	cost     int            //cost
	capacity Capacity       //capacity
	from     *Vertex
	to       *Vertex
	maxCapacity Capacity
}

func NewCostEdge(cost int, capacity Capacity, from *Vertex, to *Vertex) *Edge {
	return &Edge{
		name:     GetEdgeName(from, to),
		cost:     cost,
		capacity: capacity,
		from:     from,
		to:       to,
		maxCapacity: NewIntCapacityWithCapacity(capacity)}
}

func (edge *Edge) AddCost(cost int) {
	edge.cost += cost
}

func (edge *Edge) SetCost(cost int) {
	edge.cost = cost
}

func (edge *Edge) SetCapacity(capacity Capacity) {
	edge.capacity = capacity
}

func (edge *Edge) SetFrom(from *Vertex) {
	edge.from = from
}

func (edge *Edge) SetTo(to *Vertex) {
	edge.to = to
}

func (edge *Edge) GetCost() int {
	return edge.cost
}

func (edge *Edge) GetName() string {
	return edge.name
}

func (edge *Edge) GetCapacity() Capacity {
	return edge.capacity
}

func (edge *Edge) GetReverseCapacity() Capacity {

	reverseCapacity, error := edge.maxCapacity.Sub2(edge.capacity)

	if error != nil{
		panic("reverseCapacity can't below zero")
	}
	return reverseCapacity
}

func (edge *Edge) GetMaxCapacity() Capacity {
	return edge.maxCapacity
}

func (edge *Edge) GetFrom() *Vertex {
	return edge.from
}

func (edge *Edge) GetTo() *Vertex {
	return edge.to
}


/************************************
 *
 *    Capacity cores
 *
 ************************************/

 type Capacity interface {
	 Less(capacity Capacity) bool
	 Add(capacity Capacity)
	 Sub(capacity Capacity) error
	 Sub2(capacity Capacity) (Capacity, error)
	 IsNull() bool
	 GoString() string
 }

type IntCapacity struct {
	value int
}

func NewIntCapacity(value int) Capacity{
	return &IntCapacity{value:value}
}

func NewIntCapacityWithCapacity(cap Capacity) Capacity{
	intCapacity := cap.(*IntCapacity)
	return &IntCapacity{value:intCapacity.value}
}

func (cap *IntCapacity) Less(capacity Capacity) bool {
	intCapacity := capacity.(*IntCapacity)
	if  cap.value < intCapacity.value{
		return true
	}else {
		return false
	}
}
func (cap *IntCapacity) Add(capacity Capacity) {
	intCapacity := capacity.(*IntCapacity)
	cap.value += intCapacity.value
	return
}
func (cap *IntCapacity) Sub(capacity Capacity) error{
	intCapacity := capacity.(*IntCapacity)
	cap.value = cap.value - intCapacity.value
	if cap.value < 0 {
		cap.value += intCapacity.value
		return errors.New("capacity can't less 0")

	}
	return nil
}

func (cap *IntCapacity) Sub2(capacity Capacity) (Capacity, error){
	intCapacity := capacity.(*IntCapacity)
	value := cap.value - intCapacity.value
	if value >= 0{
		return NewIntCapacity(value),nil
	}else {
		return nil, errors.New("capacity can't less 0")
	}

}

func (cap *IntCapacity) IsNull() bool{
	if cap.value <= 0 {
		return true
	}
	return false
}
func (cap *IntCapacity) GoString() string {
	goString := "Capacity : ["+strconv.Itoa(cap.value)+"]"
	return goString
}