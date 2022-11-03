package benor

import (
	"math/rand"
	"time"
)

type Node struct {
	N, F  int
	Nodes map[int]*Node
	ID    int

	Echo1Msgs, Echo2Msgs map[int]map[bool][]int // maps round > value > node ids
	Echo2BotMsgs         map[int][]int

	Malicious bool

	Value   bool
	Round   int
	Decided bool
}

func newNode(n, f, id int, malicious bool) *Node {
	return &Node{
		N:         n,
		F:         f,
		ID:        id,
		Malicious: malicious,

		Echo1Msgs:    map[int]map[bool][]int{},
		Echo2Msgs:    map[int]map[bool][]int{},
		Echo2BotMsgs: map[int][]int{},
		Round:        0,
	}
}

func (n *Node) Start(round int, value bool) {
	if n.Malicious {
		return
	}

	if n.Decided {
		return
	}

	// prevents re-starting the same round over and over again
	if round <= n.Round {
		return
	}

	n.Round = round
	n.Value = value
	for _, node := range n.Nodes {
		node.Echo1(n.ID, n.Round, value)
	}
}

func (n *Node) Echo1(id int, round int, value bool) {
	if n.Malicious {
		return
	}

	if n.Echo1Msgs[round] == nil {
		n.Echo1Msgs[round] = map[bool][]int{}
	}
	if n.Echo1Msgs[round][value] == nil {
		n.Echo1Msgs[round][value] = []int{}
	}
	n.Echo1Msgs[round][value] = append(n.Echo1Msgs[round][value], id)

	echo1MsgCnt := len(n.Echo1Msgs[round][false]) + len(n.Echo1Msgs[round][true])

	if echo1MsgCnt >= n.N-n.F {
		if len(n.Echo1Msgs[round][false]) > (n.N+n.F)/2 {
			for _, node := range n.Nodes {
				node.Echo2D(n.ID, round, false)
			}
		} else if len(n.Echo1Msgs[round][true]) > (n.N+n.F)/2 {
			for _, node := range n.Nodes {
				node.Echo2D(n.ID, round, true)
			}
		} else {
			for _, node := range n.Nodes {
				node.Echo2Bot(n.ID, round)
			}
		}
	}
}

func (n *Node) Echo2D(id int, round int, value bool) {
	if n.Malicious {
		return
	}

	if n.Echo2Msgs[round] == nil {
		n.Echo2Msgs[round] = map[bool][]int{}
	}
	if n.Echo2Msgs[round][value] == nil {
		n.Echo2Msgs[round][value] = []int{}
	}
	n.Echo2Msgs[round][value] = append(n.Echo2Msgs[round][value], id)

	n.echo2Process(round)
}

func (n *Node) Echo2Bot(id int, round int) {
	if n.Malicious {
		return
	}

	if n.Echo2BotMsgs[round] == nil {
		n.Echo2BotMsgs[round] = []int{}
	}
	n.Echo2BotMsgs[round] = append(n.Echo2BotMsgs[round], id)

	n.echo2Process(round)
}

func (n *Node) echo2Process(round int) {
	if n.cntEcho2Msgs(round) >= n.N-n.F {
		if len(n.Echo2Msgs[round][false]) >= n.F+1 {
			n.Value = false
		}
		if len(n.Echo2Msgs[round][true]) >= n.F+1 {
			n.Value = true
		}

		if n.cntEcho2DecidedMsgs(round) >= (n.N+n.F)/2 {
			n.Decided = true
			return
		} else {
			n.Value = RandBool()
		}

		n.Start(n.Round+1, n.Value)
	}
}

func (n *Node) cntEcho2Msgs(round int) int {
	return len(n.Echo2BotMsgs[round]) + n.cntEcho2DecidedMsgs(round)
}

func (n *Node) cntEcho2DecidedMsgs(round int) int {
	return len(n.Echo2Msgs[round][true]) + len(n.Echo2Msgs[round][false])
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
