package consensus

import (
	"consensus-demo/types"
)

type BlockChain interface {
	CurrentBlock() *types.Block
}

type Consensus interface {
	NewGenesisBlock() *types.Block
	GenerateBlock(bc BlockChain, difficulty uint, body string) (*types.Block, error)
	VerifyBlock(bc BlockChain, block *types.Block) error
}

