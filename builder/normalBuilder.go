package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *NormalBuilder) setNumFloor() {
	n.floor = 2
}

func (n *NormalBuilder) getHouse() House {
	return House{
		doorType:   n.doorType,
		windowType: n.windowType,
		floor:      n.floor,
	}
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}
