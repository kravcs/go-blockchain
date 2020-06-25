package blockchain

import "math/big"

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// create the hash to see  if it meets a set of requirements

// Requirements
// The First few bytes must contain 0's

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}
