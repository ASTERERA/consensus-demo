package pow

import (
	"errors"
	"github.com/ASTERERA/consensus-demo/consensus"
	"github.com/ASTERERA/consensus-demo/types"
)

type Pow struct {
}

func NewPow() *Pow {
	return &Pow{}
}

// NewGenesisBlock 生成一个创世区块
func (p *Pow) NewGenesisBlock() *types.Block {
	block := &types.Block{}
	block.Hash = block.ToHash()

	return block
}

// GenerateBlock 产生新的区块
// block的hash值，需满足`verifyHash()函数`返回true
func (p *Pow) GenerateBlock(bc consensus.BlockChain, difficulty uint, body string) (*types.Block, error) {
	// TODO. 添加代码
	// 提示：
	// 1. 新生成的block数据：number=当前block的number+1，timestamp=当前时间，parenthash=当前block的hash值，difficulty和body为传入的值
	// 2. 使用for语句循环，每次为block.nonce随机一个值，然后使用`block.ToHash()`计算block.hash值，当hash值满足`verifyHash()`函数返回true，说明出块成功，返回新的block
	// 3. `verifyHash()`需要自己填写代码，算法见此函数注释


	return nil, errors.New("not implement yet")
}

// verifyHash 验证hash值的有效性
// 算法：首先，去掉hash值前2位的0x; 然后，验证hash值的前difficulty个字符串是否为0. 如：difficulty=2，hash值头2个字符串为0
func (p *Pow) verifyHash(hash string, difficulty uint) bool {
	var verify bool
	// TODO. 添加代码
	// 提示：按照上面注释的算法，设置verify变量的值
	return verify
}

// VerifyBlock 验证区块是否合法
func (p *Pow) VerifyBlock(bc consensus.BlockChain, block *types.Block) error {
	// TODO. 添加代码
	// 提示：
	// 1. 验证block的基础数据的正确性，字段非0，非空等：number，parentHash，hash
	// 2. 验证block的hash值是否满足挖矿难度
	// 3. 验证block和父区块(链上的当前区块)的关联是否正确（number和parentHash）


	return errors.New("not implement yet")
}