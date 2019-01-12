package merkletree

import (
	"encoding/json"
	"fmt"
)

// Node represents a single node in a Merkle tree
type Node interface {
	fmt.Stringer
	Hash() string
	Index() int
}

// MerkleTree defines and represents the methods a generic Merkle tree should have
type MerkleTree interface {
	fmt.Stringer
	Add(data []byte) (index int, hash string)
	IntermediaryHashesByIndex(index int) (intermediaryHashes []string, err error)
	ValidateExistence(original []byte, index int, intermediaryHashes []string) (bool, error)
	HashAt(index int) (string, error)
	Root() string
	Length() int
}

type internaler interface {
	Insert(hash string) (index int)
	RawInsert(hash string) (index int, leaf Node)
	Recalculate() (root string)
}

// InternalMerkleTree defines additional functions that are not supposed to be exposed to outside user to call.
// These functions deal with direct inserts of hashes and tree recalculation
type InternalMerkleTree interface {
	MerkleTree
	internaler
}

type externaler interface {
	json.Marshaler
}

// ExternalMerkleTree defines additional functions that are to be exported when the tree is communicated with the outside world.
type ExternalMerkleTree interface {
	MerkleTree
	externaler
}

// FullMerkleTree is both Internal and External
type FullMerkleTree interface {
	MerkleTree
	internaler
	externaler
}
