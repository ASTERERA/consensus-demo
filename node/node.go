package node

import (
	"errors"
	"fmt"
	"github.com/ASTERERA/consensus-demo/blockchain"
	"github.com/ASTERERA/consensus-demo/consensus"
	"github.com/ASTERERA/consensus-demo/consensus/pos"
	"github.com/ASTERERA/consensus-demo/consensus/pow"
	"github.com/ASTERERA/consensus-demo/miner"
	"github.com/ASTERERA/consensus-demo/server"
	"time"
)

type Config struct {
	Name       string
	Difficulty uint
	Consensus  string
}

type Node struct {
	conf  Config
	miner *miner.Miner
	svr   *server.Server
}

// New 创建区块链节点
func New(conf Config) (*Node, error) {
	var cs consensus.Consensus
	var chain *blockchain.BlockChain
	switch conf.Consensus {
	case "pow":
		cs = pow.NewPow()
	case "pos":
		cs = pos.NewPos()
	default:
		return nil, errors.New("wrong consensus setting, must be pow or pos")
	}
	chain = blockchain.New(cs)
	miner := miner.New(cs, conf.Difficulty, chain)
	return &Node{
		conf:  conf,
		miner: miner,
		svr:   server.New(chain, miner),
	}, nil
}

// Start 启动节点
func (n *Node) Start() {
	n.miner.Start()
	fmt.Printf("node[%s] started\n", n.conf.Name)

	n.svr.Run()
}

// Stop 停止节点
func (n *Node) Stop() {
	n.miner.Exit()
	time.Sleep(time.Second)
	fmt.Printf("node[%s] stopped\n", n.conf.Name)
}
