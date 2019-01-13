package cores

import (
	"testing"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-3
 *
 **************************************************************************************************************/

func TestIsNullString(t *testing.T) {
	if !IsNullString("") {
		t.Error()
	}
}

func TestIsNullVertex(t *testing.T) {
	var vertex *Vertex = nil
	if !IsNullVertex(vertex) {
		t.Error()
	}
}

func TestIsNullEdge(t *testing.T) {
	var edge *Edge = nil
	if !IsNullEdge(edge) {
		t.Error()
	}
}

func TestGetEdgeName(t *testing.T) {
	var s *Vertex = NewVertex("start")
	var e *Vertex = NewVertex("end")
	if GetEdgeName(s, e) != "start-end" {
		t.Error()
	}
}

