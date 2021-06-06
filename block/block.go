package block

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Hash [32]byte

type Block struct {
	Timestamp time.Time
	Hash      Hash
	PrevHash  Hash
}

func CalculateHash(prevHash Hash, timestamp time.Time) Hash {
	format := fmt.Sprintf("%s%s", prevHash, timestamp.Format(time.RFC3339))
	return sha256.Sum256([]byte(format))
}

func New(prevHash Hash, timestamp time.Time) Block {
	return Block{
		Timestamp: timestamp,
		Hash:      CalculateHash(prevHash, timestamp),
		PrevHash:  prevHash,
	}
}

func (block *Block) Dump() {
	fmt.Println("{")
	fmt.Print("    Timestamp : ")
	fmt.Println(block.Timestamp.Format(time.ANSIC))
	fmt.Print("    Hash      : ")
	fmt.Println(block.Hash)
	fmt.Print("    PrevHash  : ")
	fmt.Println(block.PrevHash[:])
	fmt.Println("}")
}

func Genisis() Block {
	return Block{
		Timestamp: time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC),
		Hash:      Hash{},
		PrevHash:  Hash{},
	}
}
