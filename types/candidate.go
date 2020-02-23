package types

import "encoding/json"

type Candidate struct {
	Name  string `json:"name"`
	Votes uint64 `json:"votes"`
}

type Candidates []*Candidate

func (cs Candidates) ToBytes() []byte {
	if cs == nil {
		return nil
	}
	bytes, _ := json.Marshal(cs)
	return bytes
}

func (cs Candidates) Clone() Candidates {
	newCs := make([]*Candidate, len(cs))
	for i, c := range cs {
		newC := &Candidate{c.Name, c.Votes}
		newCs[i] = newC
	}
	return newCs
}

// VotesSum 所有candidate的票数综合
func (cs Candidates) VotesSum() uint64 {
	var sum uint64
	for _, c := range cs {
		sum += c.Votes
	}
	return sum
}

func (cs Candidates) Equal(target Candidates) bool {
	if target == nil || len(cs) != len(target) {
		return false
	}
	for i, c := range cs {
		if c.Name != target[i].Name || c.Votes != target[i].Votes {
			return false
		}
	}
	return true
}
