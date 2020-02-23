package consensus

import (
	"errors"
	"github.com/ASTERERA/consensus-demo/types"
)

func TestFakeBlock(bc BlockChain, block *types.Block, cs Consensus, verifyBlock func(cs Consensus, bc BlockChain, block *types.Block) error) error {
	// fake the data
	// parent := bc.CurrentBlock()
	fakeBlock := block.Clone()
	fakeBlock.Difficulty = 0 // avoid the difficulty hash check

	// test block number
	fakeBlock.Number = block.Number + 1
	fakeBlock.Hash = fakeBlock.ToHash()
	if err := verifyBlock(cs, bc, fakeBlock); err == nil {
		return errors.New("it must cause an error because the number data is wrong")
	}
	// rollback
	fakeBlock.Number = block.Number
	fakeBlock.Hash = fakeBlock.ToHash()

	// test block parentHash
	fakeBlock.ParentHash = block.ParentHash[len(block.ParentHash)-1:] + block.ParentHash[1:len(block.ParentHash)-1]
	fakeBlock.Hash = fakeBlock.ToHash()
	if err := verifyBlock(cs, bc, fakeBlock); err == nil {
		return errors.New("it must cause an error because the parentHash data is wrong")
	}
	// rollback
	fakeBlock.ParentHash = block.ParentHash
	fakeBlock.Hash = fakeBlock.ToHash()

	// test the block hash
	fakeBlock.Hash = block.Hash[len(block.Hash)-1:] + block.Hash[1:len(block.Hash)-1]
	fakeBlock.Difficulty = block.Difficulty
	if err := verifyBlock(cs, bc, fakeBlock); err == nil {
		return errors.New("it must cause an error because the hash data is wrong")
	}
	// rollback
	fakeBlock.Hash = block.Hash
	fakeBlock.Hash = fakeBlock.ToHash()

	return nil
}