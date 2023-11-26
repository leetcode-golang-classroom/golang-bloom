package bloom

import (
	"fmt"
	"hash"
	"math"
	"sync"
)

type BloomFilter struct {
	bitSet []bool        // The bit array represented as a slice of bool
	m      uint64        // The number of bits in the bit set
	hashes []hash.Hash64 // The hash functions to use
	k      uint64        // The number of hash functions to uses
	mutex  sync.Mutex    // Mutex to ensure thread safety
}

func NewBloomFilter(n uint64, p float64) (*BloomFilter, error) {
	return NewBloomFilterWithHasher(n, p, NewMurMur3Hasher())
}
func NewBloomFilterWithHasher(n uint64, p float64, h Hasher) (*BloomFilter, error) {
	if n == 0 {
		return nil, fmt.Errorf("number of elements cannot be 0")
	}
	if p <= 0 || p >= 1 {
		return nil, fmt.Errorf("false positive rate must between 0 and 1")
	}
	if h == nil {
		return nil, fmt.Errorf("hasher cannot be nil")
	}
	m, k := getOptimalParams(n, p)
	return &BloomFilter{
		m:      m,
		k:      k,
		bitSet: make([]bool, m),
		hashes: h.GetHashes(k),
	}, nil
}

func getOptimalParams(n uint64, p float64) (uint64, uint64) {
	m := uint64(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	if m == 0 {
		m = 1
	}
	k := uint64(math.Ceil(float64(m)/float64(n)) * math.Log(2))
	if k == 0 {
		k = 1
	}
	return m, k
}

func (bf *BloomFilter) Add(data []byte) {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()
	for _, hash := range bf.hashes {
		hash.Reset()
		hash.Write(data)
		hashValue := hash.Sum64() % bf.m
		bf.bitSet[hashValue] = true
	}
}

func (bf *BloomFilter) Test(data []byte) bool {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()
	for _, hash := range bf.hashes {
		hash.Reset()
		hash.Write(data)
		hashValue := hash.Sum64() % bf.m
		if !bf.bitSet[hashValue] {
			return false
		}
	}
	return true
}
