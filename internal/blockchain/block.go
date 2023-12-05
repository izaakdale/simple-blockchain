package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

var genesis = newBlock([]byte("I am Genesis"), []byte{})

type block struct {
	timestamp         int64
	previousBlockHash []byte
	myBlockHash       []byte
	allData           []byte
}

func newBlock(data, prevBlockHash []byte) *block {
	b := block{
		timestamp:         time.Now().Unix(),
		previousBlockHash: prevBlockHash,
		allData:           data,
	}
	b.setHash()
	return &b
}

func (b *block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.previousBlockHash, b.allData}, []byte{})
	hash := sha256.Sum256(headers)
	b.myBlockHash = hash[:]
}
