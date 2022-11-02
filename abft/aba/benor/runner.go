package benor

// Runner is the implementation for Ben Or's Binary Agreement algo http://homepage.cs.uiowa.edu/~ghosh/BenOr.pdf
type Runner struct {
	Nodes map[int]*Node
}

func NewMoreThanFMaliciousBroadcaster() *Runner {
	return new(map[int]bool{1: true, 2: true, 3: true, 4: false, 5: false})
}

func NewLessThanFMaliciousBroadcaster() *Runner {
	return new(map[int]bool{1: true, 2: false, 3: false, 4: false, 5: false})
}

func New() *Runner {
	return new(map[int]bool{1: false, 2: false, 3: false, 4: false, 5: false})
}

func new(maliciousNodes map[int]bool) *Runner {
	n1 := newNode(5, 2, 1, maliciousNodes[1])
	n2 := newNode(5, 2, 2, maliciousNodes[2])
	n3 := newNode(5, 2, 3, maliciousNodes[3])
	n4 := newNode(5, 2, 4, maliciousNodes[4])
	n5 := newNode(5, 2, 5, maliciousNodes[5])

	n1.Nodes = map[int]*Node{
		2: n2,
		3: n3,
		4: n4,
		5: n5,
	}
	n2.Nodes = map[int]*Node{
		1: n1,
		3: n3,
		4: n4,
		5: n5,
	}
	n3.Nodes = map[int]*Node{
		1: n1,
		2: n2,
		4: n4,
		5: n5,
	}
	n4.Nodes = map[int]*Node{
		1: n1,
		2: n2,
		3: n3,
		5: n5,
	}
	n5.Nodes = map[int]*Node{
		1: n1,
		2: n2,
		3: n3,
		4: n4,
	}

	return &Runner{
		Nodes: map[int]*Node{
			1: n1,
			2: n2,
			3: n3,
			4: n4,
			5: n5,
		},
	}
}

// Run receives reliably received messages and returns output
func (r *Runner) Run(values map[int]bool) (bool, bool) {
	for id, n := range r.Nodes {
		n.Start(values[id])
	}

	return r.Nodes[5].Decided, r.Nodes[5].Value
}
