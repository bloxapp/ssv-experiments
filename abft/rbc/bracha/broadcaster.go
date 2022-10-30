package bracha

type Broadcaster struct {
	Nodes map[int]*Node
}

func NewBroadcaster() *Broadcaster {
	return newBroadcaster(map[int]bool{1: false, 2: false, 3: false, 4: false})
}

func NewLessThanFMaliciousBroadcaster() *Broadcaster {
	return newBroadcaster(map[int]bool{1: false, 2: true, 3: false, 4: false})
}

func NewMoreThanFMaliciousBroadcaster() *Broadcaster {
	return newBroadcaster(map[int]bool{1: true, 2: true, 3: false, 4: false})
}

func newBroadcaster(maliciousNodes map[int]bool) *Broadcaster {
	n1 := newNode(4, 1, 1, maliciousNodes[1])
	n2 := newNode(4, 1, 2, maliciousNodes[2])
	n3 := newNode(4, 1, 3, maliciousNodes[3])
	n4 := newNode(4, 1, 4, maliciousNodes[4])

	n1.Nodes = map[int]*Node{
		2: n2,
		3: n3,
		4: n4,
	}
	n2.Nodes = map[int]*Node{
		1: n1,
		3: n3,
		4: n4,
	}
	n3.Nodes = map[int]*Node{
		1: n1,
		2: n2,
		4: n4,
	}
	n4.Nodes = map[int]*Node{
		1: n1,
		2: n2,
		3: n3,
	}

	return &Broadcaster{
		Nodes: map[int]*Node{
			1: n1,
			2: n2,
			3: n3,
			4: n4,
		},
	}
}

func (b *Broadcaster) Broadcast(leader int, data []byte) []byte {
	for _, n := range b.Nodes {
		n.FromLeader(leader, data)
	}

	return b.Nodes[leader].DeliveredValue
}
