package test

import (
	"fmt"
	"testing"

	"github.com/leetcode-golang-classroom/golang-bloom/bloom"
)

func TestBloomFilter_AddAndTest(t *testing.T) {
	t.Parallel()
	n := uint64(10000)
	p := 0.01
	bf, err := bloom.NewBloomFilter(n, p)
	if err != nil {
		t.Errorf("Failed to create Bloom filter: %s", err)
	}

	item := "test-item"
	bf.Add([]byte(item))
	if !bf.Test([]byte(item)) {
		t.Errorf("Item '%s' should be present in the bloom filter, but it's not.", item)
	}
}

func TestBloomFilter_AddTestNonExistentItem(t *testing.T) {
	t.Parallel()
	n := uint64(1000)
	p := 0.01
	bf, err := bloom.NewBloomFilter(n, p)
	if err != nil {
		t.Errorf("Failed to create Bloom filter: %s", err)
	}
	if bf.Test([]byte("non-existent-item")) {
		t.Errorf("Non-existent item should not be present in the Bloom filter")
	}
}

func TestBloomFilter_FalsePositiveRate(t *testing.T) {
	t.Parallel()
	n := uint64(1000000)
	p := 0.00100
	bf, err := bloom.NewBloomFilter(n, p)
	if err != nil {
		t.Errorf("Failed to create Bloom filter: %s", err)
	}
	// Add n items
	for i := uint64(0); i < n; i++ {
		item := fmt.Sprintf("test-item-%d", i)
		bf.Add([]byte(item))
	}
	// Check for a different set of n items to estimate the false positive rate
	falsePositives := 0
	for i := uint64(0); i < n; i++ {
		item := fmt.Sprintf("different-item-%d", i)
		if bf.Test([]byte(item)) {
			falsePositives++
		}
	}
	estimatedFalsePositiveRate := float64(falsePositives) / float64(n)
	if estimatedFalsePositiveRate > (p * 1.15) {
		t.Errorf("Estimated false positive rate is higher than expected: got %.3f%% (%d%d), want <= %.3f%%", estimatedFalsePositiveRate*100, falsePositives, n, p*100)
	}
}
