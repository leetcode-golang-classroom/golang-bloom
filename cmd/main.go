package main

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-bloom/bloom"
)

func main() {
	bf, _ := bloom.NewBloomFilter(1000, 0.01)
	bf.Add([]byte("foo"))
	bf.Add([]byte("bar"))
	bf.Add([]byte("baz"))

	log.Println(bf.Test([]byte("foo")))
	log.Println(bf.Test([]byte("bar")))
	log.Println(bf.Test([]byte("baz")))
	log.Println(bf.Test([]byte("qux")))
}
