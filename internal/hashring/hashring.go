package hashring

import (
	"fmt"
	"hash/crc32"
	"hash/fnv"
	"sort"
)

func hashKeyCRC32(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func hashKeyFnv(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func hashKey(key string) uint32 {
    // return hashKeyCRC32(key)
    return hashKeyFnv(key)
}

func (r *HashRing) Rebalance() {
	// Calculate the greatest common divisor of the weights
	weightGcd := 0
	for _, node := range r.nodes {
		weightGcd = gcd(weightGcd, node.Weight)
	}

	// Build a Consistent Hashing Ring
	r.ring = make([]ringNode, 0)
	for i := 0; i < len(r.nodes); i++ {
		node := &r.nodes[i]

		node.virtualNodeCount = r.virtualNodePerNode * node.Weight / weightGcd

		for j := 0; j < node.virtualNodeCount; j++ {
			r.ring = append(r.ring, ringNode{
				Id:   node.Id,
				Hash: hashKey(fmt.Sprintf("%s#%d", node.Id, j)) % r.ringLength,
			})
		}

		r.ring = append(r.ring, ringNode{
			Id:   node.Id,
			Hash: hashKey(node.Id) % r.ringLength,
		})
	}
	sort.Slice(r.ring, func(i, j int) bool {
		return r.ring[i].Hash < r.ring[j].Hash
	})
}

func (r HashRing) Get(objId string) []string {
	result := make([]string, r.replicas)
	for i := 0; i < r.replicas; i++ {
		subObjId := fmt.Sprintf("%s|%d", objId, i)
		result[i] = r.get(subObjId)
	}
	return result
}

func (r HashRing) get(objId string) string {
	objHash := hashKey(objId) % r.ringLength

	left, right := 0, len(r.ring)-1
	var ans int

	for left <= right {
		mid := (left + right) / 2
		if r.ring[mid].Hash < objHash {
			ans = mid + 1
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}
	if ans >= len(r.ring) {
		ans = len(r.ring) - 1
	}

	return r.ring[ans].Id
}
