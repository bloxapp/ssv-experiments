package bracha

// Node https://decentralizedthoughts.github.io/2020-09-19-living-with-asynchrony-brachas-reliable-broadcast/
type Node struct {
	N, F               int
	Nodes              map[int]*Node
	EchoMsgs, VoteMsgs map[int][]byte
	DeliveredValue     []byte

	Malicious bool

	id         int
	echo, vote bool
}

func newNode(n, f, id int, malicious bool) *Node {
	return &Node{
		Malicious: malicious,
		N:         n,
		F:         f,
		EchoMsgs:  map[int][]byte{},
		VoteMsgs:  map[int][]byte{},
		id:        id,
		echo:      true,
		vote:      true,
	}
}

func (b *Node) FromLeader(fromID int, data []byte) {
	if b.Malicious {
		return
	}
	if b.echo {
		b.ReceiveEcho(b.id, data)
		for _, n := range b.Nodes {
			n.ReceiveEcho(b.id, data)
		}
		b.echo = false
	}
}

func (b *Node) ReceiveEcho(fromID int, data []byte) {
	b.EchoMsgs[fromID] = data
	if len(b.EchoMsgs) >= b.N-b.F && b.vote {
		b.ReceiveVote(b.id, data)
		for _, n := range b.Nodes {
			n.ReceiveVote(b.id, data)
		}
		b.vote = false
	}
}

func (b *Node) ReceiveVote(fromID int, data []byte) {
	b.VoteMsgs[fromID] = data
	//if len(b.VoteMsgs) >= b.F+1 && b.vote {
	//	b.ReceiveVote(b.id, data)
	//	for _, n := range b.Nodes {
	//		n.ReceiveVote(b.id, data)
	//	}
	//	b.vote = false
	//}
	if len(b.VoteMsgs) >= b.N-b.F {
		b.DeliveredValue = data
	}
}
