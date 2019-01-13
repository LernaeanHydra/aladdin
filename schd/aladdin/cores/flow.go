package cores

import (
	"log"
	//"fmt"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-3
 *
 **************************************************************************************************************/

type Flow struct {
	capacity Capacity
	paths    []Path
}

func NewFlow() Flow {
	return Flow{capacity: nil,
		paths: []Path{}}
}

func (flow *Flow) AddFlow(capacity Capacity) {
	//fmt.Println("------"+capacity.GoString())
	if flow.capacity == nil {
		//fmt.Println("flow capacity is nil")
		flow.capacity = NewIntCapacityWithCapacity(capacity)
		//fmt.Println(flow.GetFlow().GoString())
	} else {
		flow.capacity.Add(capacity)
	}
}

func (flow *Flow) GetFlow() Capacity {
	return flow.capacity
}

func (flow *Flow) AddPath(path Path) {
	if &path == nil || path.cost == 0 {
		log.Print("Invalid path")
		return
	}
	flow.paths = append(flow.paths, path)
}

func (flow *Flow) GetPaths() []Path {
	return flow.paths
}


