package solvers

import (
	"k8s.io/kubernetes/schd/aladdin/cores"
	"fmt"
	//"k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/storage/value"
	// "github.com/onsi/gomega/matchers/support/goraph/edge"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-3
 *
 **************************************************************************************************************/

 type ShortestPathPolicy interface {
	Push(v cores.Vertex)
	Pop() cores.Vertex
	Len() int
	Next(b bool, m map[string]string, v cores.Vertex)
 }

 type Solver struct {
	graph  *cores.Graph
	source cores.Vertex
	sink   cores.Vertex
	policy ShortestPathPolicy
 }

 func NewShortestPathSolver(graph *cores.Graph,
 				source cores.Vertex,
 				sink cores.Vertex,
 				policy ShortestPathPolicy) Solver {
	return Solver{
		graph: graph,
		source: source,
		sink: sink,
		policy: policy,
	}
 }

func NewSMaxFlowSolver(graph *cores.Graph,
	source cores.Vertex,
	sink cores.Vertex,
	policy ShortestPathPolicy) Solver {
	return Solver{
		graph: graph,
		source: source,
		sink: sink,
		policy: policy,
	}
}

 func (s *Solver) ShortestPath() cores.Path {
	 allPaths  := make(map[string]cores.Path)
	 visited   := make(map[string]string)
	 s.policy.Push(s.source)

	 for {
		 if s.policy.Len() == 0 { // queue has no vertex
			 break
		 }

		 allPaths[cores.GetEdgeName(&s.source,
			 &s.source)] = *cores.NewPath()   // initial source to source path as 0 cost

		 vertex := s.policy.Pop()
		 name := vertex.GetName()
		 visited[name] = name

		 fmt.Println("<!--------"+vertex.GetName()+"-------->")
		 for _, edge := range vertex.GetOutEdges() {
			 fmt.Println("<!--------"+edge.GoString()+"-------->")
			 pathName := cores.GetEdgeName(&s.source, edge.GetTo())

			 pathCost := -1 // infinity
			 if path, ok := allPaths[pathName]; ok {
				 pathCost = path.GetCost()
			 }

			 priorPath := allPaths[cores.GetEdgeName(
				 &s.source, &vertex)]
			 priorCost := priorPath.GetCost()

			 currentPath := cores.NewPath()

			 // if cost need to update
			 if !edge.GetCapacity().IsNull() && (pathCost == -1 || pathCost > priorCost + edge.GetCost()) {
				 fmt.Println("<!--------need to update-------->")
			 	// clone prior path to current path
				 if priorPath.GetEdges() != nil {
					 for edge2 := priorPath.GetEdges().Front(); edge2 != nil; edge2 = edge2.Next() {

						 value := edge2.Value.(cores.Edge)

						 clone := cores.NewCostEdge(value.GetCost(),cores.NewIntCapacityWithCapacity(value.GetCapacity()),
							 value.GetFrom(), value.GetTo())
						 fmt.Print("******"+clone.GoString())
						 currentPath.AddEdge(*clone)
					 }
				 }
				 // add new edge to current path
				 clone_edge := cores.NewCostEdge(edge.GetCost(),cores.NewIntCapacityWithCapacity(edge.GetCapacity()), edge.GetFrom(), edge.GetTo())

				 currentPath.AddEdge(*clone_edge)
				 // update cost
				 fmt.Println("<!--------pathName "+pathName +":"+currentPath.GoString()+"-------->")
				 allPaths[pathName] = *currentPath
				 edge.GetTo().SetDistance(priorCost + edge.GetCost())
			 }


			 // whether toVertex is visited before, if false then push it to queue
			 if !edge.GetCapacity().IsNull() {
				 _, ok := visited[edge.GetTo().GetName()]
				 s.policy.Next(ok, visited, *edge.GetTo())
			 }
		 }
	 }

	 // if there is a path from source to sink, then return it. or return -1 empty edge
	 if _, ok := allPaths[cores.GetEdgeName(&s.source, &s.sink)]; ok {
	 	fmt.Println("start print allPaths:")
	 	for key, value := range allPaths {
	 		fmt.Println(key)
	 		fmt.Println(value.GoString())
		}
		 fmt.Println("end print allPaths:")
	 	fmt.Println("***************"+allPaths[cores.GetEdgeName(&s.source, &s.sink)].GoString())
		 return	allPaths[cores.GetEdgeName(&s.source, &s.sink)]
	 } else {
	 	return *cores.NewInvalidPath()
	 }
 }

func (s *Solver) AvailablePath() cores.Path {
	for _, firstOutEdge := range s.source.GetOutEdges(){
		// 找到所有未调度的任务节点
		if firstOutEdge.GetCapacity().IsNull() {
			continue
		}
		firstLevelVertex := firstOutEdge.GetTo()
		for secondEdgeName, secondOutEdge := range firstLevelVertex.GetOutEdges(){
			// 找到所有的机器节点
			secondLevelVertex := secondOutEdge.GetTo()
			// 是否可以直接部署到该机器
			capacityNeed := cores.NewIntCapacity(0)  // 如果不能部署，需要反悔的最少容量；最大容量是它本身
			for _, toSinkEdge := range secondLevelVertex.GetOutEdges(){
				// 如果可以直接部署
				if !toSinkEdge.GetCapacity().Less(firstOutEdge.GetCapacity()) {
					currentPath := cores.NewPath()
					currentPath.AddEdge(firstOutEdge)
					currentPath.AddEdge(secondOutEdge)
					currentPath.AddEdge(toSinkEdge)
					return *currentPath
				}else {

					tmpCapacity, err := firstOutEdge.GetCapacity().Sub2(toSinkEdge.GetCapacity())
					if err != nil {
						panic("this step shouldn't excute")
					}
					capacityNeed.Add(tmpCapacity)
					fmt.Println("schedule pod "+firstLevelVertex.GetName()+" to node "+secondLevelVertex.GetName()+" need other "+capacityNeed.GoString())
				}
			}
			// 不能直接部署，就需要考虑反悔哪一个任务
			var selectedEdge cores.Edge
			var selectedEdgeName = ""

			for thirdEdgeName, thirdInEdge := range secondLevelVertex.GetInEdges(){
				// 反悔任务时本任务，跳过
				if thirdEdgeName == secondEdgeName{
					continue
				}
				// 反悔任务流量小于最小容量需求 或者 不小于任务本身的容量需求，跳过

				fmt.Println("---开始获得反向边容量---")
				fmt.Println("Edge's maxCapacity : "+thirdInEdge.GetMaxCapacity().GoString())
				fmt.Println("Edge's capacity : "+thirdInEdge.GetCapacity().GoString())
				fmt.Println("---结束获得反向边容量---")

				fmt.Println("Selecting a task to prrmpt: this edge "+thirdInEdge.GetName()+"'s reverseEdge has "+thirdInEdge.GetReverseCapacity().GoString())
				fmt.Println("Selecting a task to prrmpt: this pod needs "+firstOutEdge.GetCapacity().GoString())

				if thirdInEdge.GetReverseCapacity().Less(capacityNeed) || !thirdInEdge.GetReverseCapacity().Less(firstOutEdge.GetCapacity()){
					continue
				}
				// 找到大于最小容量需求的最小任务进行反悔
				if selectedEdgeName == "" {
					selectedEdge = thirdInEdge
					selectedEdgeName = thirdEdgeName
					continue
				}
				if thirdInEdge.GetReverseCapacity().Less(selectedEdge.GetReverseCapacity()) {
					selectedEdge = thirdInEdge
					selectedEdgeName = thirdEdgeName
				}
				
			}
			// 如果找不到可反悔的任务，则直接返回无可行路径
			if selectedEdgeName == ""{
				fmt.Println("Can't find a task to preempt")
				break
			}

			thirdLevelVertex := selectedEdge.GetFrom()
			// 反悔的任务是否可以放在其他的机器上
			for fourthEdgeName, fourthOutEdge := range thirdLevelVertex.GetOutEdges(){
				if fourthEdgeName == selectedEdgeName {
					continue
				}
				fourthLevelVertex := fourthOutEdge.GetTo() // 获得待反悔的机器节点
				// 判断该待反悔的机器节点是否有足够的容量
				for _, toSinkEdge := range fourthLevelVertex.GetOutEdges(){
					// 如果可以反悔到这个机器上
					if !toSinkEdge.GetCapacity().Less(fourthOutEdge.GetCapacity()) {
						currentPath := cores.NewPath()
						currentPath.AddEdge(firstOutEdge)
						currentPath.AddEdge(secondOutEdge)
						currentPath.AddEdge(selectedEdge)
						currentPath.AddEdge(fourthOutEdge)
						currentPath.AddEdge(toSinkEdge)
						return *currentPath
					}
				}

			}
			// 如果不能反悔到别的机器上，就直接反悔整个任务，返回的任务不部署
			for _, fromSourceEdge := range thirdLevelVertex.GetInEdges(){
				currentPath := cores.NewPath()
				currentPath.AddEdge(firstOutEdge)
				currentPath.AddEdge(secondOutEdge)
				currentPath.AddEdge(selectedEdge)
				currentPath.AddEdge(fromSourceEdge)
				return *currentPath
			}


		}
	}

	return *cores.NewInvalidPath()
}

 func (s *Solver) MaxFlow() cores.Flow {
	//s.graph.PrintGragh()
 	//degs := make(map[string]string)
	flow := cores.NewFlow()

	for {
		// get current shortest path
		// path := s.ShortestPath()

		// get available path by adding reverse arc
		path := s.AvailablePath()

		// if path's cost is -1, then means no path from source to sink
		if path.GetCost() == -1 {
			break;
		}

		minFlow := s.graph.UpdateGraghForMaxFlow(path)

		flow.AddFlow(minFlow)
		flow.AddPath(path)
		s.graph.PrintGragh()
	}
	return flow
  }