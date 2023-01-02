package hashring

type Node struct {
	Id               string
	Weight           int
	virtualNodeCount int
}

type ringNode struct {
	Id     string
	Hash   uint32
	Length uint32
}

type HashRing struct {
	nodes              []Node
	virtualNodePerNode int
	ring               []ringNode
	ringLength         uint32
	replicas           int
}

func New(nodes []Node, replicas int) *HashRing {
	return &HashRing{
		nodes:              nodes,
		virtualNodePerNode: 100000,
		ring:               make([]ringNode, 0),
		ringLength:         268435456,
		replicas:           replicas,
	}
}
