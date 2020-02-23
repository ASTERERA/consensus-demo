package miner

import (
	"consensus-demo/blockchain"
	"consensus-demo/consensus"
	"consensus-demo/consensus/pow"
	"consensus-demo/transation"
	"consensus-demo/types"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"
)

const (
	MineInterval = 2
)

type Miner struct {
	cs         consensus.Consensus
	difficulty uint
	working    int32
	exit       chan struct{}
	chain      *blockchain.BlockChain
	TxManager  *transation.TxManager
}

// New 创建miner
func New(cs consensus.Consensus, difficulty uint, chain *blockchain.BlockChain) *Miner {
	m := &Miner{cs: cs, difficulty: difficulty, chain: chain, exit: make(chan struct{}), TxManager: transation.NewTxManager()}
	// start work loop
	go m.Work()
	return m
}

// Start 开始挖矿
func (m *Miner) Start() {
	if !m.Working() {
		atomic.StoreInt32(&m.working, 1)
	}
	fmt.Println("miner started")
}

// Stop 停止挖矿
func (m *Miner) Stop() {
	atomic.StoreInt32(&m.working, 0)
	fmt.Println("miner stopped")
}

// Exit 程序退出
func (m *Miner) Exit() {
	close(m.exit)
}

// Working 是否正在挖矿
func (m *Miner) Working() bool {
	return atomic.LoadInt32(&m.working) == 1
}

// Work 挖矿调度函数
func (m *Miner) Work() {
	ticker := time.NewTicker(MineInterval * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if !m.Working() {
				continue
			}
			var block *types.Block
			var err error
			if _, ok := m.cs.(*pow.Pow); ok { // pow
				block, err = m.cs.GenerateBlock(m.chain, m.difficulty, "")
				if err != nil {
					fmt.Printf("failed to generate new block: + %s\n", err.Error())
					continue
				}
				fmt.Printf("block[#%d] was mined. hash: %s\n", block.Number, block.Hash)
			} else { // pos
				txs := m.TxManager.FetchTxsThenClear()
				// mine block only when income new transactions
				// if len(txs) <= 0 {
				// 	continue
				// }
				bytes, _ := json.Marshal(txs)
				block, err = m.cs.GenerateBlock(m.chain, m.difficulty, string(bytes))
				if err != nil {
					fmt.Printf("failed to generate new block: + %s\n", err.Error())
					// push back the txs into the tx pool
					m.TxManager.AddTxs(txs)
					continue
				}
				fmt.Printf("block[#%d] was mined. validator: %s, hash: %s\n", block.Number, block.Validator, block.Hash)
			}

			if err := m.chain.AddBlock(block); err != nil {
				fmt.Printf("failed to verify block: %s\n", err.Error())
			}
		case <-m.exit:
			fmt.Println("miner exited")
			return
		}
	}
}
