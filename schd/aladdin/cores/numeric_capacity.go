package cores

import (
	"strconv"
	"errors"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-3
 *
 **************************************************************************************************************/

type NumericCapacity struct {
	capacity int
}

func NewNumericCapacity(capacity int) NumericCapacity {
	return NumericCapacity{capacity: capacity}
}

// please check the parameters (capacity) by yourself
func (nc NumericCapacity) Less(capacity Capacity) bool {
 	oc := capacity.(NumericCapacity)
	return nc.capacity < oc.capacity
}

// please check the parameters (capacity) by yourself
func (nc NumericCapacity) Add(capacity Capacity) {
	oc := capacity.(NumericCapacity)
	nc.capacity += oc.capacity
}

// please check the parameters (capacity) by yourself
func (nc NumericCapacity) Sub(capacity Capacity) error {
	intCapacity := capacity.(*IntCapacity)
	nc.capacity -= intCapacity.value
	if nc.capacity < 0 {
		return errors.New("capacity can't less 0")
	}
	return nil
}

// please check the parameters (capacity) by yourself
func (nc NumericCapacity) Sub2(capacity Capacity) (Capacity, error) {

	return nil,nil
}

func (sc NumericCapacity) IsNull() bool {
	return sc.capacity == 0
}

func (sc NumericCapacity) GoString() string {
	return strconv.Itoa(sc.capacity)
}