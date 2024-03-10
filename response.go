package idgenmodel

type IDGenResp struct {
	ID       []byte `json:"id"`
	BitCount uint   `json:"bit-count"`
}
