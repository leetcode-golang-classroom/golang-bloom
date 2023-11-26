package bloom

import (
	"hash"

	"github.com/spaolacci/murmur3"
)

type MurMur3Hahser struct{}

var _ Hasher = (*MurMur3Hahser)(nil)

func NewMurMur3Hasher() *MurMur3Hahser {
	return &MurMur3Hahser{}
}

func (h *MurMur3Hahser) GetHashes(n uint64) []hash.Hash64 {
	hashers := make([]hash.Hash64, n)
	for i := 0; uint64(i) < n; i++ {
		hashers[i] = murmur3.New64WithSeed(uint32(i))
	}

	return hashers
}
