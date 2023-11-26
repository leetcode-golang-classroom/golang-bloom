# golang-bloom

This repository is implementation for bloom filter with golang

## introduction 

Bloom filter is a mechanism that use hash function to map message into a bitwise map

which could help to judge whether is a message is not in the bitwise map.

However, due to the length of the bitwise are limited and hash function need to uniformly

Bloom filter is false positive, which mean even if Bloom filter return true, the message could be not included in Bloom filter Record.

Only could make sure, whether the information is not in Bloom filter record

