package benor

// Runner is the implementation for Ben Or's Binary Agreement algo http://homepage.cs.uiowa.edu/~ghosh/BenOr.pdf
type Runner struct {
	Nodes map[int]*Node
}

func NewMoreThanFMaliciousBroadcaster() *Runner {
	return new(map[int]bool{1: true, 2: true, 3: false, 4: false, 5: false, 6: false})
}

func NewLessThanFMaliciousBroadcaster() *Runner {
	return new(map[int]bool{1: true, 2: false, 3: false, 4: false, 5: false, 6: false})
}

func New() *Runner {
	return new(map[int]bool{1: false, 2: false, 3: false, 4: false, 5: false, 6: false})
}

func new(maliciousNodes map[int]bool) *Runner {
	n1 := newNode(6, 1, 1, maliciousNodes[1])
	n2 := newNode(6, 1, 2, maliciousNodes[2])
	n3 := newNode(6, 1, 3, maliciousNodes[3])
	n4 := newNode(6, 1, 4, maliciousNodes[4])
	n5 := newNode(6, 1, 5, maliciousNodes[5])
	n6 := newNode(6, 1, 5, maliciousNodes[6])

	nodes := map[int]*Node{
		1: n1,
		2: n2,
		3: n3,
		4: n4,
		5: n5,
		6: n6,
	}
	n1.Nodes = nodes
	n2.Nodes = nodes
	n3.Nodes = nodes
	n4.Nodes = nodes
	n5.Nodes = nodes
	n6.Nodes = nodes

	return &Runner{
		Nodes: map[int]*Node{
			1: n1,
			2: n2,
			3: n3,
			4: n4,
			5: n5,
			6: n6,
		},
	}
}

// Run receives reliably received messages and returns output
func (r *Runner) Run(values map[int]bool) (bool, bool) {
	for id, n := range r.Nodes {
		n.Start(1, values[id])
	}

	return r.Nodes[5].Decided, r.Nodes[5].Value
}
