package server

import (
	"consensus-demo/blockchain"
	"consensus-demo/miner"
	"consensus-demo/transation"
	"consensus-demo/types"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Server struct {
	http.Server
	chain *blockchain.BlockChain
	miner *miner.Miner
}

func New(chain *blockchain.BlockChain, m *miner.Miner) *Server {
	r := gin.Default()
	s := &Server{
		chain: chain,
		miner: m,
	}
	s.Addr = ":8080"
	s.Handler = r

	s.Router()

	return s
}

// Router 设置api请求路由
func (s *Server) Router() {
	group := s.Handler.(*gin.Engine).Group("/api")
	{
		// blockchain api
		group.GET("/blockchain/latest", s.GetLatestBlock)
		group.GET("/blockchain/height", s.GetBlockchainHeight)
		group.GET("/blockchain", s.GetBlockChain)

		// miner api
		group.GET("/miner/stop", s.StopMine)
		group.GET("/miner/start", s.StartMine)

		// block api
		group.GET("/block/:number", s.GetBlockNumber)

		// pos api
		group.GET("/pos/vote", s.PosVote)
		group.GET("/pos/pendingTxs", s.PosPendingTxs)
	}
}

func (s *Server) Run() {
	s.ListenAndServe()
}

func (s *Server) Stop() {
	s.Shutdown(nil)
}

// GetLatestBlock 获取最新区块
func (s *Server) GetLatestBlock(c *gin.Context) {
	block := s.chain.CurrentBlock()
	c.JSON(200, block)
	// http.StatusOK, gin.H{
}

// GetBlockchainHeight 获取当前区块高度
func (s *Server) GetBlockchainHeight(c *gin.Context) {
	block := s.chain.CurrentBlock().Number
	c.JSON(200, block)
	// http.StatusOK, gin.H{
}

// StopMine 停止挖矿
func (s *Server) StopMine(c *gin.Context) {
	s.miner.Stop()
}

// StartMine 开始挖矿
func (s *Server) StartMine(c *gin.Context) {
	s.miner.Start()
}

// GetBlockNumber 通过区块高度获取区块
func (s *Server) GetBlockNumber(c *gin.Context) {
	strNum := c.Param("number")
	number, _ := strconv.ParseUint(strNum, 10, 64)
	block := s.chain.GetBlock(number)
	c.JSON(200, block)
}

// GetBlockChain 获取整条区块链数据
func (s *Server) GetBlockChain(c *gin.Context) {
	c.JSON(200, s.chain.GetBlockChain())
}

// GetBlockChain 用户投票
func (s *Server) PosVote(c *gin.Context) {
	candidate := c.Query("candidate")
	if candidate == "" {
		c.JSON(400, "lack param: candidate, please set a value by c1/c2/c3")
		return
	}
	votesStr := c.Query("votes")
	if votesStr == "" {
		c.JSON(400, "lack param: votes, please set a int value")
		return
	}
	votes, err := strconv.ParseInt(votesStr, 10, 64)
	if err != nil {
		c.JSON(400, "wrong param: votes, please set a int value")
		return
	}

	cv := &types.Candidate{Name: candidate, Votes: uint64(votes)}
	bytes, _ := json.Marshal(cv)
	s.miner.TxManager.AddTx(transation.Tx(bytes))

	c.JSON(200, "success")
}

// PosPending 获取当前尚未打包的交易信息
func (s *Server) PosPendingTxs(c *gin.Context) {
	c.JSON(200, s.miner.TxManager.PendingTxs())
}

