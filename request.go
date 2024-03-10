package idgenmodel

import "github.com/prakash-p-3121/errorlib"

type IDGenReq struct {
	TableName *string
}

func (req *IDGenReq) Validate() errorlib.AppError {
	if req.TableName == nil {
		return errorlib.NewBadReqError("table-name-null")
	}
	return nil
}
