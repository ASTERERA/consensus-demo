package pos

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ASTERERA/consensus-demo/consensus"
	"github.com/ASTERERA/consensus-demo/transation"
	"github.com/ASTERERA/consensus-demo/types"
)

// 3个默认的候选者
var (
	DefaultCandidates = []string{"c1", "c2", "c3"}
)

type Pos struct {
}

func NewPos() *Pos {
	return &Pos{}
}

// NewGenesisBlock 生成一个创世区块
func (p *Pos) NewGenesisBlock() *types.Block {
	// init candidates
	candidatesMap := make(map[string]uint64)
	for _, candidate := range DefaultCandidates {
		candidatesMap[candidate] = 0
	}
	block := &types.Block{}
	block.Validator = DefaultCandidates[0]
	candidates := make(types.Candidates, 0)
	for _, candidate := range DefaultCandidates {
		candidates = append(candidates, &types.Candidate{Name: candidate, Votes: 1})
	}
	block.Candidates = candidates

	bytes, _ := json.Marshal(block.Candidates)
	block.Body = string(bytes)
	block.Hash = block.ToHash()

	return block
}

// GenerateBlock 产生新的区块
func (p *Pos) GenerateBlock(bc consensus.BlockChain, difficulty uint, body string) (*types.Block, error) {
	// TODO. 添加代码
	// 提示：
	// 1. 新生成的block数据：number=当前block的number+1，timestamp=当前时间，parenthash=当前block的hash值，difficulty和body为传入的值
	// 2. body中保存的是所有投票数据，转换成`[]Tx`之后，通过调用函数`addVotesToCandidates()`获取最新的candidates
	// 3. 随机一个nonce值，然后通过`findValidator()`获取validator值
	// 4. 最后别忘了设置Hash值


	return nil, errors.New("not implement yet")
}

// addVotesToCandidates 为candidates添加票数, 返回新的candidates实例
func (p *Pos) addVotesToCandidates(oldCandidates types.Candidates, txs []transation.Tx) types.Candidates {
	newCandidates := oldCandidates.Clone()
	// TODO. 添加代码


	return newCandidates
}

// findValidator 通过index值查找validator
// 算法：从第一个candidate开始，累加票数，当index值小于此票数时，取当前candidate作为validator
func (p *Pos) findValidator(candidates types.Candidates, index uint64) string {
	// TODO. 添加代码

	return ""
}

// verifyCandidatesAndValidator 验证candidates和validator
func (p *Pos) verifyCandidatesAndValidator(parent *types.Block, block *types.Block) error {
	var txs []transation.Tx
	if block.Body != "" {
		if err := json.Unmarshal([]byte(block.Body), &txs); err != nil {
			return err
		}
	}
	candidates := p.addVotesToCandidates(parent.Candidates, txs)
	if !candidates.Equal(block.Candidates) {
		return errors.New("wrong candidates data")
	}
	validator := p.findValidator(candidates, block.Nonce)
	if block.Validator != validator {
		return fmt.Errorf("wrong validator, expect: %s, get: %s", block.Validator, validator)
	}

	return nil
}

// VerifyBlock 验证区块是否合法
func (p *Pos) VerifyBlock(bc consensus.BlockChain, block *types.Block) error {
	// TODO. 添加代码
	// 提示：
	// 1. 验证block的基础数据的正确性，字段非0，非空等：number，parentHash，hash
	// 2. 验证block和父区块(链上的当前区块)的关联是否正确（number和parentHash）
	// 3. 验证candidates和validator数据是否正确



	return nil
}
