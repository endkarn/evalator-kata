package evalator

const loadlimit = 1000

type Lift struct {
	Weight       int
	CurrentFloor int
}

func (lift Lift) CheckLoadLimit() bool {
	if lift.Weight > loadlimit {
		return false
	}
	return true
}
