package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"reflect"
	"strconv"
	"time"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	myBlock := Block{"My Block", time.Now().Unix(), []byte{}, []byte{}}
	myBlock.Hash = myBlock.calculateHash()
	Blockchain = []Block{myBlock}
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	newBlock := Block{data, time.Now().Unix(), []byte{}, []byte{}}
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = newBlock.calculateHash()
	return newBlock
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	lastBlock := Blockchain[len(Blockchain)-1]
	if !reflect.DeepEqual(lastBlock.Hash, b.PrevHash) {
		return errors.New("Invalid block !")
	}
	Blockchain = append(Blockchain, b)
	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
