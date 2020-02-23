package transation

import "sync"

type Tx string

type TxManager struct {
	txs []Tx
	mu  sync.RWMutex
}

func NewTxManager() *TxManager {
	return &TxManager{
		txs: make([]Tx, 0),
	}
}

func (tm *TxManager) PendingTxs() []Tx {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	return tm.txs
}

// FetchTxsThenClear 获取所有txs并清空txs记录
func (tm *TxManager) FetchTxsThenClear() []Tx {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	txs := tm.txs
	tm.txs = make([]Tx, 0)

	return txs
}

func (tm *TxManager) AddTx(tx Tx) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.txs = append(tm.txs, tx)
	return nil
}

func (tm *TxManager) AddTxs(txs []Tx) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.txs = append(tm.txs, txs...)
	return nil
}