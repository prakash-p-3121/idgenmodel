package idgenmodel

type IDGenResp struct {
	ID       []byte `json:"id"`
	BitCount int64  `json:"bit-count"`
}
