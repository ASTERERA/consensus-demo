package blockchain

import (
	"consensus-demo/consensus"
	"consensus-demo/types"
	"sync"
)

type BlockChain struct {
	blocks  []*types.Block
	cs      consensus.Consensus
	blockMu sync.RWMutex
	bcMu    sync.Mutex
}

// New 创建一条新链
func New(cs consensus.Consensus) *BlockChain {
	blocks := make([]*types.Block, 1)
	// init genesis block
	blocks[0] = cs.NewGenesisBlock()

	chain := &BlockChain{blocks: blocks, cs: cs}

	return chain
}

// AddBlock 添加新的block到链上
func (bc *BlockChain) AddBlock(block *types.Block) error {
	bc.bcMu.Lock()
	defer bc.bcMu.Unlock()
	if err := bc.cs.VerifyBlock(bc, block); err != nil {
		return err
	}
	bc.blockMu.Lock()
	defer bc.blockMu.Unlock()
	bc.blocks = append(bc.blocks, block)
	return nil
}

// GetBlock 获取指定高度的block
func (bc *BlockChain) GetBlock(number uint64) *types.Block {
	if number > bc.CurrentBlock().Number {
		return nil
	}

	bc.blockMu.RLock()
	defer bc.blockMu.RUnlock()

	return bc.blocks[number]
}

// GetBlockChain 获取整条区块链数据
func (bc *BlockChain) GetBlockChain() []*types.Block {
	bc.blockMu.RLock()
	defer bc.blockMu.RUnlock()

	return bc.blocks
}

// CurrentBlock 获取当前高度的block
func (bc *BlockChain) CurrentBlock() *types.Block {
	bc.blockMu.RLock()
	defer bc.blockMu.RUnlock()

	l := len(bc.blocks)
	return bc.blocks[l-1]
}

// // PendingVotes 尚未被打包的投票结果
// func (bc *GetBlockChain) PendingVotes() *types.Block {
//
// }
//
// // PendingVotes 尚未被打包的投票结果
// func (bc *GetBlockChain) PendingVotes() *types.Block {
//
// }
