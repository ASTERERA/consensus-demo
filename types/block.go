package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/ASTERERA/consensus-demo/utils"
)

type Block struct {
	Number     uint64
	Timestamp  uint64
	Hash       string
	ParentHash string
	Difficulty uint
	Nonce      uint64 // 在pow算法中，表示满足Difficulty的随机值；在pos算法中，表示validator所在位置（Candidates数组下标index）
	Body       string // 在pos算法中，保存对应的candidate的投票数, json格式

	// used for `pos` consensus to record validators
	Validator  string
	Candidates Candidates // 保存每个candidate的票数
}

// ToHash 使用sha256算法，生成hash值
func (b *Block) ToHash() string {
	hash := sha256.New()
	hash.Write(b.toBytes())
	return "0x" + hex.EncodeToString(hash.Sum(nil))
}

// toBytes 拼接除了`Hash`字段外的其他所有
func (b *Block) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.Write(utils.Int64ToBytes(int64(b.Number)))
	buffer.Write(utils.Int64ToBytes(int64(b.Timestamp)))
	buffer.Write([]byte(b.ParentHash))
	buffer.Write(utils.IntToBytes(int(b.Difficulty)))
	buffer.Write(utils.Int64ToBytes(int64(b.Nonce)))
	buffer.Write([]byte(b.Body))

	if b.Validator != "" && b.Candidates != nil { // data for pos
		buffer.Write([]byte(b.Validator))
		buffer.Write(b.Candidates.ToBytes())
	}

	return buffer.Bytes()
}

func (b *Block) Clone() *Block {
	bytes, _ := json.Marshal(b)
	block := new(Block)
	_ = json.Unmarshal(bytes, block)
	return block
}
