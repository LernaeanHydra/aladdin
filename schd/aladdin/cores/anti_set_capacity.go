package cores


/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-22
 *
 **************************************************************************************************************/

type AntiSetCapacity struct {
	name        string
	deployed    map[string]int
	constraints *AntiAffinity
}

func NewAntiSetCapacity(name string, aa *AntiAffinity) *AntiSetCapacity {
	return &AntiSetCapacity{
		name:        name,
		deployed:    make(map[string]int),
		constraints: aa,
	}
}

// please check the parameters (capacity) by yourself
func (nc AntiSetCapacity) Less(capacity Capacity) bool {
	oc := capacity.(AntiSetCapacity)

	_, exist := nc.deployed[oc.name]

	if exist == false {
		return true
	}

	name := AntiAffinityKey(nc.name, oc.name)
	value := nc.constraints.GetAntiAffinity()[name]

	if value > 0 {
		if _, ok := nc.deployed[oc.name]; !ok || nc.deployed[oc.name] + 1 <= value {
			return true
		}
		return false
	} else if value == 0 {
		return false
	} else {
		if _, ok := nc.deployed[nc.name]; !ok  || nc.deployed[nc.name] + 1 <= -value {
			return true
		}
		return false
	}

}

// please check the parameters (capacity) by yourself
func (nc AntiSetCapacity) Add(capacity Capacity) {
	oc := capacity.(AntiSetCapacity)

	_, exist := nc.deployed[oc.name]

	if exist == false {
		nc.deployed[oc.name] = 1
	} else {
		nc.deployed[oc.name]++
	}
}

// please check the parameters (capacity) by yourself
func (nc AntiSetCapacity) Sub(capacity Capacity) error {

	oc := capacity.(AntiSetCapacity)

	_, exist := nc.deployed[oc.name]

	if exist == true {
		nc.deployed[oc.name]--
		if nc.deployed[oc.name] == 0 {
			delete(nc.deployed, oc.name)
		}
	}

	return nil
}

func (nc AntiSetCapacity) Sub2(capacity Capacity)(Capacity,error) {

	return nil, nil
}

func (sc AntiSetCapacity) IsNull() bool {
	return true
}

func (sc AntiSetCapacity) GoString() string {
	return sc.name
}