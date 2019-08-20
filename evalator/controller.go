package evalator

const loadlimit = 1000

type Lift struct {
	Weight       int
	CurrentFloor int
}

func (lift Lift) checkLoadLimit() bool {
	if lift.Weight > loadlimit {
		return false
	}
	return true
}

func (lift Lift) checkIndicator(destinationFloor int) string {
	return "Up"
}
