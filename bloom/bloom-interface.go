package bloom

import "hash"

type BloomInterface interface {
	Add([]byte)
	Test([]byte) bool
}

type Hasher interface {
	GetHashes(n uint64) []hash.Hash64
}
