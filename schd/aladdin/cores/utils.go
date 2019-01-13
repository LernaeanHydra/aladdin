package cores

import (
	"strconv"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-2
 *
 **************************************************************************************************************/


func IsNullVertex (vertex *Vertex) bool {
	if vertex == nil {
		return true
	}
	return false
}

func IsNullEdge (edge *Edge) bool {
	if edge == nil {
		return true
	}
	return false
}

func IsVertexExist (vertices map[string]Vertex, key string) bool {
	if _, ok := vertices[key]; ok {
		return true
	}
	return false
}

func IsEdgeExist (edges map[string]Edge, key string) bool {
	if _, ok := edges[key]; ok {
		return true
	}
	return false
}

func IsNullString (str string) bool {
	if str == "" {
		return true
	}
	return false
}

func AntiAffinityKey (str1 string, str2 string) string {
	return str1 + "---" + str2
}

func GetEdgeName(from *Vertex, to *Vertex) string {
	if from == nil || to == nil {
		return ""
	}
	return from.GetName() + "-" + to.GetName()
}

/*
 * GoString
 */

func (vertex *Vertex) GoString() string {
	str := "Name=" + vertex.name + "\n"
	for _, edge := range vertex.outEdges {
		str += "\t" + edge.GoString() + "\n"
	}
	return str
}

func (edge Edge) GoString() string {
	str := "[" + edge.from.name + "] -> [" + edge.to.name + "] = " + strconv.Itoa(edge.cost)
	return str
}

func (path Path) GoString() string {
	str := "path's edge number is "+strconv.Itoa(path.GetEdges().Len()) +"\n"
	for edge := path.edges.Front(); edge != nil; edge = edge.Next() {
		str += "\t" + edge.Value.(Edge).GoString()+ "\n"
	}
	return str
}

func (flow Flow) GoString() string {
	str := "Flow = " + flow.capacity.GoString() + "\n"
	for _, path := range flow.paths {
		str += "\tCost = " + strconv.Itoa(path.cost) + "\n"
		for edge := path.edges.Front(); edge != nil; edge = edge.Next() {
			str += "\t\t" + edge.Value.(Edge).GoString() + "\n"
		}
	}
	return str
}
